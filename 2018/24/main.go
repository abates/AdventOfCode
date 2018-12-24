package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type AttackType int

const (
	Bludgeoning AttackType = iota
	Fire
	Slashing
	Radiation
	Cold
)

func (at *AttackType) UnmarshalText(text []byte) (err error) {
	str := string(text)
	switch str {
	case "bludgeoning":
		*at = Bludgeoning
	case "fire":
		*at = Fire
	case "slashing":
		*at = Slashing
	case "radiation":
		*at = Radiation
	case "cold":
		*at = Cold
	default:
		err = fmt.Errorf("Unknown attack type %q", str)
	}
	return err
}

func (at AttackType) String() (str string) {
	switch at {
	case Bludgeoning:
		str = "Bludgeoning"
	case Fire:
		str = "Fire"
	case Slashing:
		str = "Slashing"
	case Radiation:
		str = "Radiation"
	case Cold:
		str = "Cold"
	default:
		str = "Unknown"
	}
	return str
}

type AttackTypeIndex map[AttackType]struct{}

func (atl *AttackTypeIndex) Contains(attackType AttackType) bool {
	_, found := (*atl)[attackType]
	return found
}

func (atl *AttackTypeIndex) UnmarshalText(text []byte) (err error) {
	for _, token := range bytes.Split(text, []byte(", ")) {
		var at AttackType
		err = at.UnmarshalText(token)
		if err != nil {
			break
		}
		(*atl)[at] = struct{}{}
	}
	return err
}

type GroupType int

const (
	ImmuneSystem GroupType = iota
	Infection
)

func (gt GroupType) String() string {
	if gt == ImmuneSystem {
		return "Immune System"
	} else if gt == Infection {
		return "Infection"
	}
	return "Unknown"
}

type AttackGroup struct {
	id         int
	groupType  GroupType
	units      int
	hitPoints  int
	damage     int
	initiative int

	weaknesses AttackTypeIndex
	immunities AttackTypeIndex
	attackType AttackType

	currentOpponent *AttackGroup
}

func (ag *AttackGroup) Select(enemies *GroupList) int {
	selected := -1
	maxDamage := 0
	ag.currentOpponent = nil
	for i, enemy := range enemies.groups {
		if enemy.groupType == ag.groupType {
			continue
		}

		damage := enemy.EffectiveDamage(ag.units*ag.damage, ag.attackType)
		if damage <= 0 {
			continue
		}

		if ag.currentOpponent == nil || maxDamage < damage {
			maxDamage = damage
			selected = i
			ag.currentOpponent = enemy
		} else if maxDamage == damage {
			if ag.currentOpponent.EffectivePower() < enemy.EffectivePower() {
				selected = i
				ag.currentOpponent = enemy
			} else if ag.currentOpponent.EffectivePower() == enemy.EffectivePower() {
				if ag.currentOpponent.initiative < enemy.initiative {
					selected = i
					ag.currentOpponent = enemy
				}
			}
		}
	}
	return selected
}

func (ag *AttackGroup) Attack() (damage int) {
	if ag.currentOpponent != nil {
		damage = ag.currentOpponent.EffectiveDamage(ag.units*ag.damage, ag.attackType)
		delta := (damage / ag.currentOpponent.hitPoints)
		if delta > ag.currentOpponent.units {
			delta = ag.currentOpponent.units
		}
		ag.currentOpponent.units -= delta
	}
	return
}

func (ag *AttackGroup) EffectiveDamage(damage int, attackType AttackType) int {
	if ag.immunities.Contains(attackType) {
		return 0
	}

	if ag.weaknesses.Contains(attackType) {
		return damage * 2
	}
	return damage
}

func (ag *AttackGroup) EffectivePower() int {
	return ag.units * ag.damage
}

func (ag *AttackGroup) UnmarshalText(text []byte) (err error) {
	ag.weaknesses = make(AttackTypeIndex)
	ag.immunities = make(AttackTypeIndex)
	line := make([]byte, len(text))
	copy(line, text)
	if o := bytes.Index(line, []byte("(")); o >= 0 {
		if c := bytes.Index(line, []byte(")")); c >= 0 {
			tokens := bytes.Split(line[o+1:c], []byte("; "))
			for _, token := range tokens {
				if i := bytes.Index(token, []byte("weak to ")); i >= 0 {
					err = ag.weaknesses.UnmarshalText(token[len("weak to "):])
				} else if i := bytes.Index(token, []byte("immune to ")); i >= 0 {
					err = ag.immunities.UnmarshalText(token[len("immune to "):])
				} else {
					err = fmt.Errorf("Could not parse %q", string(token))
				}
				if err != nil {
					break
				}
			}
			if err == nil {
				line = append(line[0:o], line[c+1:]...)
			}
		} else {
			err = fmt.Errorf("Open parenthesis but no close parenthesis")
		}
	}

	if err == nil {
		attackType := ""
		_, err = fmt.Sscanf(string(line), "%d units each with %d hit points with an attack that does %d %s damage at initiative %d", &ag.units, &ag.hitPoints, &ag.damage, &attackType, &ag.initiative)
		if err == nil {
			err = ag.attackType.UnmarshalText([]byte(attackType))
		}
	}
	return err
}

type GroupList struct {
	groups []*AttackGroup
}

func (gl *GroupList) Append(group *AttackGroup) {
	gl.groups = append(gl.groups, group)
}

func (gl *GroupList) Len() int               { return len(gl.groups) }
func (gl *GroupList) Remove(i int)           { gl.groups = append(gl.groups[0:i], gl.groups[i+1:]...) }
func (gl *GroupList) Get(i int) *AttackGroup { return gl.groups[i] }

func (gl *GroupList) TotalUnits() (units int) {
	for _, group := range gl.groups {
		units += group.units
	}
	return units
}

func (gl *GroupList) Copy() (groups *GroupList) {
	groups = &GroupList{}
	for _, group := range gl.groups {
		groups.Append(group)
	}
	return groups
}

func (gl *GroupList) Filter(groupType GroupType) (groups *GroupList) {
	groups = &GroupList{}
	for _, group := range gl.groups {
		if group.groupType == groupType {
			groups.Append(group)
		}
	}
	return groups
}

func (gl *GroupList) String() string {
	var builder strings.Builder
	if len(gl.groups) == 0 {
		builder.WriteString("No groups remain.\n")
	} else {
		for _, group := range gl.groups {
			builder.WriteString(fmt.Sprintf("Group %d contains %d units ep: %d\n", group.id, group.units, group.EffectivePower()))
		}
	}
	return builder.String()
}

func (gl *GroupList) Swap(i, j int) {
	tmp := gl.groups[i]
	gl.groups[i] = gl.groups[j]
	gl.groups[j] = tmp
}

type attack struct {
	*GroupList
}

func AttackOrder(gl *GroupList) sort.Interface {
	return &attack{gl}
}

func (at *attack) Less(i, j int) bool {
	return at.GroupList.Get(i).initiative < at.GroupList.Get(j).initiative
}

type selection struct {
	*GroupList
}

func SelectionOrder(gl *GroupList) sort.Interface {
	return &selection{gl}
}

func (sel *selection) Less(i, j int) bool {
	ag := sel.GroupList.Get(i)
	other := sel.GroupList.Get(j)
	if ag.EffectivePower() < other.EffectivePower() {
		return true
	}

	if ag.EffectivePower() == other.EffectivePower() {
		return ag.initiative < other.initiative
	}

	return false
}

type Battle struct {
	groups *GroupList
}

func (b *Battle) UnmarshalText(text []byte) (err error) {
	var groupType GroupType
	b.groups = &GroupList{}
	id := 1
	for _, line := range bytes.Split(text, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		if bytes.HasPrefix(line, []byte("Immune")) {
			groupType = ImmuneSystem
			continue
		}

		if bytes.HasPrefix(line, []byte("Infection")) {
			id = 1
			groupType = Infection
			continue
		}
		group := &AttackGroup{id: id, groupType: groupType}
		id++
		err = group.UnmarshalText(line)
		if err == nil {
			b.groups.Append(group)
		} else {
			break
		}
	}
	return err
}

func (b *Battle) ImmuneSystem() *GroupList { return b.groups.Filter(ImmuneSystem) }
func (b *Battle) Infection() *GroupList    { return b.groups.Filter(Infection) }

func (b *Battle) Fight() int {
	targets := b.groups.Copy()

	// target selection
	sort.Sort(sort.Reverse(SelectionOrder(b.groups)))
	for _, group := range b.groups.groups {
		i := group.Select(targets)
		if i >= 0 {
			//fmt.Printf("%v group %d selected target %+v\n", group.groupType, group.id, group.currentOpponent)
			targets.Remove(i)
		}
	}

	//fmt.Println()

	// attack
	attackCount := 0
	sort.Sort(sort.Reverse(AttackOrder(b.groups)))
	for _, group := range b.groups.groups {
		if group.currentOpponent != nil {
			ep := group.currentOpponent.EffectivePower()
			group.Attack()
			if group.currentOpponent.EffectivePower() < ep {
				attackCount++
			}
		}
	}

	// remove groups with no remaining units
	for i := 0; i < b.groups.Len(); i++ {
		if b.groups.Get(i).EffectivePower() <= 0 {
			b.groups.Remove(i)
			i--
		}
	}
	return attackCount
}

func (b *Battle) Battle(immuneBoost int) (immuneSystem, infection *GroupList) {
	immuneSystem = b.ImmuneSystem()
	for _, group := range immuneSystem.groups {
		group.damage += immuneBoost
	}

	infection = b.Infection()
	for b.ImmuneSystem().Len() > 0 && b.Infection().Len() > 0 {
		attackCount := b.Fight()
		immuneSystem = b.ImmuneSystem()
		infection = b.Infection()
		if attackCount == 0 {
			break
		}
	}
	return immuneSystem, infection
}

func part2(input []byte) (err error) {
	for boost := 0; ; boost++ {
		fmt.Printf("Checking boost %d\r", boost)
		battle := &Battle{}
		err = battle.UnmarshalText(input)
		if err != nil {
			break
		}
		immuneSystem, infection := battle.Battle(boost)
		if infection.Len() == 0 {
			fmt.Printf("\nPart 2: Immune system wins with %d units\n", immuneSystem.TotalUnits())
			break
		}
	}
	return err
}

func part1(input []byte) error {
	battle := &Battle{}
	err := battle.UnmarshalText(input)
	immuneSystem, infection := battle.Battle(0)
	if infection.Len() == 0 {
		fmt.Printf("Part 1: Immune system wins with %d units\n", immuneSystem.TotalUnits())
	} else {
		fmt.Printf("Part 1: Infection wins with %d units\n", infection.TotalUnits())
	}
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}

	for i, f := range []func([]byte) error{part1, part2} {
		err = f(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i, err)
			os.Exit(-1)
		}
	}
}
