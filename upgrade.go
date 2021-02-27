package upgrade

import "sort"

type (
	Upgrade interface {
		Upgrade()
		Support() bool
		Order() int
	}

	Upgrader struct {
		Upgrades []Upgrade
	}
)

func (u *Upgrader) Register(upgrades []Upgrade) {
	sort.Slice(upgrades, func(i int, j int) bool {
		return upgrades[i].Order() < upgrades[j].Order()
	})

	u.Upgrades = upgrades
}

func (u *Upgrader) Run() {
	for _, v := range u.Upgrades {
		if v.Support() {
			v.Upgrade()
		}
	}
}
