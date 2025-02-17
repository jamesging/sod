package mage

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            SoD Phase 4 Item Sets
///////////////////////////////////////////////////////////////////////////

var ItemSetSorcerersRegalia = core.NewItemSet(core.ItemSet{
	Name: "Sorcerer's Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Your spellcasts have a 6% chance to energize you for 300 mana.
		4: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450527}
			manaMetrics := c.NewManaMetrics(actionID)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				Name:       "S03 - Mana Proc on Cast - Magister's Regalia",
				Callback:   core.CallbackOnSpellHitDealt,
				Outcome:    core.OutcomeLanded,
				ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
				ProcChance: 0.06,
				Handler: func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
					if c.HasManaBar() {
						c.AddMana(sim, 300, manaMetrics)
					}
				},
			})
		},
		// +8 All Resistances.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.ArcaneResistance: 8,
				stats.FireResistance:   8,
				stats.FrostResistance:  8,
				stats.NatureResistance: 8,
				stats.ShadowResistance: 8,
			})
		},
		// +200 Armor.
		8: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Armor, 200)
		},
	},
})

var ItemSetArcanistMoment = core.NewItemSet(core.ItemSet{
	Name: "Arcanist Moment",
	Bonuses: map[int32]core.ApplyEffect{
		// Your Temporal Beacons last 20% longer.
		2: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases all chronomantic healing you deal by 10%.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// Each time you heal a target with Regeneration, the remaining cooldown on Rewind Time is reduced by 1 sec.
		6: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

var ItemSetArcanistInsight = core.NewItemSet(core.ItemSet{
	Name: "Arcanist Insight",
	Bonuses: map[int32]core.ApplyEffect{
		// You are immune to all damage while channeling Evocation.
		2: func(agent core.Agent) {
			// May important later but for now nothing to do
		},
		// You gain 1% increased damage for 15 sec each time you cast a spell from a different school of magic.
		4: func(agent core.Agent) {
			// TODO: This is all a bit of an assumption about how this may work without having more information.
			// We may need to rework it as we get more information
			mage := agent.(MageAgent).GetMage()

			damageMultiplierPerSchool := 1.01
			auraDuration := time.Second * 15

			arcaneAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Arcane)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexArcane)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			fireAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Fire)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexFire)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			frostAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Frost)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexFrost)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			holyAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Holy)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexHoly)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			natureAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Nature)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexNature)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			shadowAura := mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus (Shadow)",
				ActionID: core.ActionID{SpellID: 456398}.WithTag(int32(stats.SchoolIndexShadow)),
				Duration: auraDuration,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier *= damageMultiplierPerSchool
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					mage.PseudoStats.DamageDealtMultiplier /= damageMultiplierPerSchool
				},
			})

			mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 4P Bonus",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell.SpellSchool.Matches(core.SpellSchoolPhysical) {
						return
					}

					if spell.SpellSchool.Matches(core.SpellSchoolArcane) {
						arcaneAura.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolFire) {
						fireAura.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolFrost) {
						frostAura.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolHoly) {
						holyAura.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolNature) {
						natureAura.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolShadow) {
						shadowAura.Activate(sim)
					}
				},
			})
		},
		// Mage Armor increases your mana regeneration while casting by an additional 15%. Molten Armor increases your spell damage and healing by 18. Ice Armor grants 20% increased chance to trigger Fingers of Frost.
		6: func(agent core.Agent) {
			mage := agent.(MageAgent).GetMage()

			bonusFoFProcChance := .20
			bonusSpiritRegenRateCasting := .15
			bonusSpellPower := 18.0

			mage.RegisterAura(core.Aura{
				Label:    "S03 - Item - T1 - Mage - Damage 6P Bonus",
				ActionID: core.ActionID{SpellID: 456402},
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					if mage.IceArmorAura != nil {
						mage.FingersOfFrostProcChance += bonusFoFProcChance
					} else if mage.MageArmorAura != nil {
						mage.PseudoStats.SpiritRegenRateCasting += bonusSpiritRegenRateCasting
					} else if mage.MoltenArmorAura != nil {
						mage.AddStatDynamic(sim, stats.SpellPower, bonusSpellPower)
					}
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					if mage.IceArmorAura != nil {
						mage.FingersOfFrostProcChance -= bonusFoFProcChance
					} else if mage.MageArmorAura != nil {
						mage.PseudoStats.SpiritRegenRateCasting += bonusSpiritRegenRateCasting
					} else if mage.MoltenArmorAura != nil {
						mage.AddStatDynamic(sim, stats.SpellPower, bonusSpellPower)
					}
				},
			})
		},
	},
})
