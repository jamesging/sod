package hunter

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            SoD Phase 3 Item Sets
///////////////////////////////////////////////////////////////////////////

var ItemSetDreadHuntersChain = core.NewItemSet(core.ItemSet{
	Name: "Dread Hunter's Chain",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 20)
			c.AddStat(stats.RangedAttackPower, 20)
		},
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.MeleeCrit, 1*core.CritRatingPerCritChance)
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            SoD Phase 4 Item Sets
///////////////////////////////////////////////////////////////////////////

var ItemSetBeastmasterArmor = core.NewItemSet(core.ItemSet{
	Name: "Beastmaster Armor",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
		// Your melee and ranged autoattacks have a 6% chance to energize you for 300 mana.
		4: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450577}
			manaMetrics := c.NewManaMetrics(actionID)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID:   actionID,
				Name:       "S03 - Mana Proc on Cast - Beaststalker Armor",
				Callback:   core.CallbackOnSpellHitDealt,
				Outcome:    core.OutcomeLanded,
				ProcMask:   core.ProcMaskWhiteHit,
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

var ItemSetGiantstalkerProwess = core.NewItemSet(core.ItemSet{
	Name: "Giantstalker Prowess",
	Bonuses: map[int32]core.ApplyEffect{
		// Your Mongoose Bite also reduces its target's chance to Dodge by 1% and increases your chance to hit by 1% for 30 sec.
		2: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()

			debuffAuras := hunter.NewEnemyAuraArray(func(target *core.Unit, level int32) *core.Aura {
				return target.RegisterAura(core.Aura{
					Label:    "S03 - Item - T1 - Hunter - Melee 2P Bonus",
					ActionID: core.ActionID{SpellID: 456389},
					Duration: time.Second * 30,
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Dodge, -1)
						aura.Unit.PseudoStats.BonusMeleeHitRatingTaken += 1 * core.MeleeHitRatingPerHitChance
						aura.Unit.PseudoStats.BonusSpellHitRatingTaken += 1 * core.SpellHitRatingPerHitChance
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Dodge, 1)
						aura.Unit.PseudoStats.BonusMeleeHitRatingTaken += 1 * core.MeleeHitRatingPerHitChance
						aura.Unit.PseudoStats.BonusSpellHitRatingTaken += 1 * core.SpellHitRatingPerHitChance
					},
				})
			})

			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "S03 - Item - T1 - Hunter - Melee 2P Bonus Trigger",
				OnSpellHitDealt: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell.SpellCode == SpellCode_HunterMongooseBite && result.Landed() {
						debuffAuras.Get(result.Target).Activate(sim)
					}
				},
			}))
		},
		// While tracking a creature type, you deal 3% increased damage to that creature type.
		// Unsure if this stacks with the Pursuit 4p
		4: func(agent core.Agent) {
			c := agent.GetCharacter()
			// Just adding 3% damage to assume the hunter is tracking their target's type
			c.PseudoStats.DamageDealtMultiplier *= 1.03
		},
		// Mongoose Bite also activates for 5 sec whenever your target Parries or Blocks or when your melee attack misses.
		6: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()
			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "S03 - Item - T1 - Hunter - Melee 6P Bonus Trigger",
				OnSpellHitDealt: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell.ProcMask.Matches(core.ProcMaskMelee) && (result.Outcome == core.OutcomeMiss || result.Outcome == core.OutcomeBlock || result.Outcome == core.OutcomeParry) {
						hunter.DefensiveState.Activate(sim)
					}
				},
			}))
		},
	},
})

var ItemSetGiantstalkerPursuit = core.NewItemSet(core.ItemSet{
	Name: "Giantstalker Pursuit",
	Bonuses: map[int32]core.ApplyEffect{
		// You generate 100% more threat for 8 sec after using Distracting Shot.
		2: func(agent core.Agent) {
			// Nothing to do
		},
		// While tracking a creature type, you deal 3% increased damage to that creature type.
		// Unsure if this stacks with the Prowess 4p
		4: func(agent core.Agent) {
			c := agent.GetCharacter()
			// Just adding 3% damage to assume the hunter is tracking their target's type
			c.PseudoStats.DamageDealtMultiplier *= 1.03
		},
		// Your next Shot ability within 12 sec after Aimed Shot deals 20% more damage.
		6: func(agent core.Agent) {
			hunter := agent.(HunterAgent).GetHunter()

			if !hunter.Talents.AimedShot {
				return
			}

			procAura := hunter.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 456379},
				Label:    "S03 - Item - T1 - Hunter - Ranged 6P Bonus",
				Duration: time.Second * 12,

				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range hunter.Shots {
						if spell != nil {
							spell.DamageMultiplier *= 1.20
						}
					}
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range hunter.Shots {
						if spell != nil {
							spell.DamageMultiplier /= 1.20
						}
					}
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if !spell.Flags.Matches(SpellFlagShot) || (aura.RemainingDuration(sim) == aura.Duration && spell.SpellCode == SpellCode_HunterAimedShot) {
						return
					}

					aura.Deactivate(sim)
				},
			})

			core.MakePermanent(hunter.RegisterAura(core.Aura{
				Label: "S03 - Item - T1 - Hunter - Ranged 6P Bonus Trigger",
				OnCastComplete: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell.SpellCode == SpellCode_HunterAimedShot {
						procAura.Activate(sim)
					}
				},
			}))
		},
	},
})
