package vanilla

import (
	"fmt"
	"time"

	"github.com/wowsims/sod/sim/common/itemhelpers"
	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/proto"
	"github.com/wowsims/sod/sim/core/stats"
)

// Ordered by ID
const (
	ShortswordOfVengeance          = 754
	FieryWarAxe                    = 870
	Bloodrazor                     = 809
	HammerOfTheNorthernWind        = 810
	FlurryAxe                      = 871
	SkullflameShield               = 1168
	Nightblade                     = 1982
	Shadowblade                    = 2163
	GutRipper                      = 2164
	HandOfEdwardTheOdd             = 2243
	BowOfSearingArrows             = 2825
	Gutwrencher                    = 5616
	Ravager                        = 7717
	HanzoSword                     = 8190
	TheJackhammer                  = 9423
	PendulumOfDoom                 = 9425
	BloodletterScalpel             = 9511
	TheHandOfAntusul               = 9639
	GryphonRidersStormhammer       = 9651
	Firebreather                   = 10797
	VilerendSlicer                 = 11603
	HookfangShanker                = 11635
	LinkensSwordOfMastery          = 11902
	SearingNeedle                  = 12531
	PipsSkinner                    = 12709
	ArcaniteChampion               = 12790
	MasterworkStormhammer          = 12794
	SerpentSlicer                  = 13035
	SealOfTheDawn                  = 13209
	JoonhosMercy                   = 17054
	Deathbringer                   = 17068
	ViskagTheBloodletter           = 17075
	ThrashBlade                    = 17705
	SatyrsLash                     = 17752
	MarkOfTheChosen                = 17774
	Thunderfury                    = 19019
	EbonHand                       = 19170
	DarkmoonCardHeroism            = 19287
	DarkmoonCardBlueDragon         = 19288
	DarkmoonCardMaelstrom          = 19289
	Nightfall                      = 19169
	RuneOfTheDawn                  = 19812
	ScarabBrooch                   = 21625
	MarkOfTheChampionPhys          = 23206
	MarkOfTheChampionSpell         = 23207
	BlisteringRagehammer           = 220569 // 10626
	SulfurasHandOfRagnaros         = 227683 // 17182
	SulfuronHammer                 = 227684 // 17193
	TemperedBlackAmnesty           = 227832 // 19166
	EbonFist                       = 227842
	ReavingNightfall               = 227843
	SkyridersMasterworkStormhammer = 227886
	FlameWrath                     = 227934 // 11809
	LordGeneralsSword              = 227940 // 11817
	WraithScythe                   = 227941
	SecondWind                     = 227967 // 11819
	BurstOfKnowledge               = 227972
	HandOfInjustice                = 227990
	Ironfoe                        = 227991 // 11684
	EbonHiltOfMarduk               = 227993 // 14576
	FrightskullShaft               = 227994 // 14531
	BarovianFamilySword            = 227997 // 14541
	Frightalon                     = 228015 // 14024
	HeadmastersCharge              = 228022 // 13937
	GravestoneWarAxe               = 228029 // 13983
	FiendishMachete                = 228056 // 18310
	RefinedArcaniteChampion        = 228125
	TalismanOfEphemeralPower       = 228255 // 18820
	GutgoreRipper                  = 228267 // 17071
	Shadowstrike                   = 228272 // 17074
	Thunderstrike                  = 228273 // 17223
	BonereaversEdge                = 228288 // 17076
	BonereaversEdgeMolten          = 228461
	EssenceOfThePureFlame          = 228293 // 18815
	PerditionsBlade                = 228296 // 18816
	Typhoon                        = 228347 // 18542
	EskhandarsLeftClaw             = 228349 // 18202
	EskhandarsRightClaw            = 228350 // 18203
	BlazefuryMedallion             = 228354 // 17111
	EmpyreanDemolisher             = 228397 // 17112
	DreadbladeOfTheDestructor      = 228410
	DreadbladeOfTheDestructor2     = 228498
	PerditionsBladeMolten          = 228511
	SkullforgeReaver               = 228542 // 13361
	RunebladeOfBaronRivendare      = 228543 // 13505
	HeartOfWyrmthalak              = 228599 // 22321
	Venomspitter                   = 228573 // 13183
	SmolderwebsEye                 = 228576 // 13213
	Chillpike                      = 228586 // 13148
	FangOfTheCrystalSpider         = 228592 // 13218
	BlackhandDoomsaw               = 228603 // 12583
	BlackbladeOfShahram            = 228606 // 12592
	SeepingWillow                  = 228666 // 12969
	DraconicInfusedEmblem          = 228678 // 22268
	HandOfJustice                  = 228722 // 11815
	Felstriker                     = 228757 // 12590
)

func init() {
	core.AddEffectsToTest = false

	// ! Please keep items ordered alphabetically within a given category !

	///////////////////////////////////////////////////////////////////////////
	//                                 Weapons
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/spell=16916/strength-of-the-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(ArcaniteChampion, "Arcanite Champion", 1.0, strengthOfTheChampionAura)

	// https://www.wowhead.com/classic/item=227997/barovian-family-sword
	// Chance on hit: Deals 30 Shadow damage every 3 sec for 15 sec. All damage done is then transferred to the caster.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(BarovianFamilySword, "Barovian Family Sword", 0.5, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 18652}

		// Keep track of damage taken by each enemy
		enemyDamageTaken := map[int32]float64{}
		for _, target := range character.Env.Encounter.TargetUnits {
			enemyDamageTaken[target.UnitIndex] = 0
		}

		healthMetrics := character.NewHealthMetrics(actionID)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 5,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Siphon Health (Barovian Family Sword)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					enemyDamageTaken[target.UnitIndex] = 0
					dot.Snapshot(target, 30, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
					enemyDamageTaken[target.UnitIndex] += result.Damage
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				enemyDamageTaken[target.UnitIndex] = 0
				spell.Dot(target).Apply(sim)
			},
		})

		// The healing is applied at the end of the DoT and can crit according to old comments
		for _, dot := range spell.Dots() {
			if dot != nil {
				unit := dot.Unit
				dot.ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
					// TODO: This may not be quite correct but it's close enough
					result := spell.CalcDamage(sim, unit, enemyDamageTaken[unit.UnitIndex], spell.OutcomeHealingCrit)
					character.GainHealth(sim, result.Damage, healthMetrics)
				})
			}
		}

		return spell
	})

	// https://www.wowhead.com/classic/item=228606/blackblade-of-shahram
	// Chance on hit: Summons the infernal spirit of Shahram.
	// Summons an NPC "Shahram" who has an equal chance to cast one of 6 spells:
	// Curse of Shahram: -50% movement speed and -25% attack speed on all enemies within 10 yards of Shahram for 10 seconds.
	// Might of Shahram: 5-second stun on all enemies within 10 yards of Shahram.
	// Fist of Shahram: +30% Melee Attack Speed for all party members within 30 yards of Shahram for 8 seconds.
	// Blessing of Shahram: Restores 50 health and mana every 5 seconds for all party members within 30 yards of Shahram for 20 seconds. The Healing portion of this effect scales at 100% of self-healing buffs such as Amplify Magic.
	// Will of Shahram: +50 all stats for yourself for 20 seconds.
	// Flames of Shahram: Deals 100-150 Fire damage to all enemies within 10 yards of Shahram. Damage scales at 100% with +spelldmg debuffs placed on enemies such as Flame Buffet.
	//
	// Implementing this without the guardian as it seems to just cast a spell and depart and guardians are expensive
	core.NewItemEffect(BlackbladeOfShahram, func(agent core.Agent) {
		character := agent.GetCharacter()

		curseOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			aura := target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16597},
				Label:    "Curse of Shahram",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1/1.25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1.25)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})
		curseOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16597},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				curseOfShahramAuras.Get(target).Activate(sim)
			},
		})

		mightOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16600},
				Label:    "Might of Shahram",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = true
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = false
				},
			})
		})
		mightOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16600},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					mightOfShahramAuras.Get(aoeTarget).Activate(sim)
				}
			},
		})

		fistOfShahramAuras := character.NewPartyAuraArray(func(unit *core.Unit) *core.Aura {
			return unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16601},
				Label:    "Fist of Shahram",
				Duration: time.Second * 8,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					character.MultiplyAttackSpeed(sim, 1.3)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					character.MultiplyAttackSpeed(sim, 1/1.3)
				},
			})
		})
		fistOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16601},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aura := range fistOfShahramAuras {
					aura.Activate(sim)
				}
			},
		})

		blessingOfShahramManaMetrics := character.NewPartyManaMetrics(core.ActionID{SpellID: 16599})
		blessingOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16599},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagIgnoreAttackerModifiers,
			Hot: core.DotConfig{
				Aura: core.Aura{
					Label: "Blessing of Shahram",
				},
				NumberOfTicks: 4,
				TickLength:    time.Second * 5,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
					dot.SnapshotBaseDamage = 50
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
					if target.HasManaBar() {
						target.AddMana(sim, 50, blessingOfShahramManaMetrics[target.UnitIndex])
					}
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, agent := range character.Party.PlayersAndPets {
					spell.Hot(&agent.GetCharacter().Unit).Apply(sim)
				}
			},
		})

		willOfShahramAura := character.GetOrRegisterAura(core.Aura{
			ActionID:  core.ActionID{SpellID: 16598},
			Label:     "Will of Shahram",
			Duration:  time.Second * 20,
			MaxStacks: 5,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
				stats := stats.Stats{
					stats.Agility:   25,
					stats.Intellect: 25,
					stats.Stamina:   25,
					stats.Spirit:    25,
					stats.Strength:  25,
				}
				character.AddStatsDynamic(sim, stats.Multiply(float64(-1*oldStacks)))
				character.AddStatsDynamic(sim, stats.Multiply(float64(newStacks)))
			},
		})
		willOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16598},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				willOfShahramAura.Activate(sim)
				willOfShahramAura.AddStack(sim)
			},
		})

		flamesOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16596},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagIgnoreAttackerModifiers,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 90, spell.OutcomeMagicCrit)
				}
			},
		})

		castableSpells := []*core.Spell{curseOfShahram, mightOfShahram, fistOfShahram, blessingOfShahram, willOfShahram, flamesOfShahram}
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Summon Shahram",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				spellIdx := int32(sim.Roll(0, 6))
				castableSpells[spellIdx].Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228603/blackhand-doomsaw
	// Chance on hit: Wounds the target for 324 to 540 damage.
	// TODO: Proc rate based on the original item
	itemhelpers.CreateWeaponProcDamage(BlackhandDoomsaw, "Blackhand Doomsaw", 0.4, 16549, core.SpellSchoolPhysical, 324, 216, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=220569/blistering-ragehammer
	// Chance on hit: Increases damage done by 20 and attack speed by 5% for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(BlisteringRagehammer, "Blistering Ragehammer", 1.0, enrageAura446327)

	itemhelpers.CreateWeaponProcDamage(BloodletterScalpel, "Bloodletter Scalpel", 1.0, 18081, core.SpellSchoolPhysical, 60, 10, 0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponProcSpell(Bloodrazor, "Bloodrazor", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17504},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Bloodrazor)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 12, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=228288/bonereavers-edge
	// https://www.wowhead.com/classic/item=228461/bonereavers-edge
	// Chance on hit: Your attacks ignore 700 of your enemies' armor for 10 sec. This effect stacks up to 3 times.
	itemhelpers.CreateWeaponProcSpell(BonereaversEdge, "Bonereaver's Edge", 2.0, makeBonereaversEdgeEffect)
	itemhelpers.CreateWeaponProcSpell(BonereaversEdgeMolten, "Bonereaver's Edge (Molten)", 2.0, makeBonereaversEdgeEffect)

	itemhelpers.CreateWeaponProcSpell(BowOfSearingArrows, "Bow of Searing Arrows", 3.35, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 29638},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeRanged,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(18, 26)
				spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeRangedCritOnly)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228586/chillpike
	// Chance on hit: Blasts a target for 160 to 250 Frost damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcDamage(Chillpike, "Chillpike", 1.0, 19260, core.SpellSchoolFrost, 160, 90, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17068/deathbringer
	// Chance on hit: Sends a shadowy bolt at the enemy causing 110 to 140 Shadow damage.
	itemhelpers.CreateWeaponProcDamage(Deathbringer, "Deathbringer", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228410/dreadblade-of-the-destructor
	// https://www.wowhead.com/classic/item=228498/dreadblade-of-the-destructor
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(DreadbladeOfTheDestructor, "Dreadblade of the Destructor", 1.0, makeDreadbladeOfTheDestructorEffect)
	itemhelpers.CreateWeaponProcSpell(DreadbladeOfTheDestructor2, "Dreadblade of the Destructor", 1.0, makeDreadbladeOfTheDestructorEffect)

	// https://www.wowhead.com/classic/item=227842/ebon-fist
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcDamage(EbonFist, "Ebon Fist", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19170/ebon-hand
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcDamage(EbonHand, "Ebon Hand", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=227993/ebon-hilt-of-marduk
	// Chance on hit: Corrupts the target, causing 210 damage over 3 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(EbonHiltOfMarduk, "Ebon Hilt of Marduk", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18656},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Corruption (Ebon Hilt of Marduk)",
				},
				TickLength:    time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 70, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228397/empyrean-demolisher
	// Chance on hit: Increases your attack speed by 20% for 10 sec.
	itemhelpers.CreateWeaponProcAura(EmpyreanDemolisher, "Empyrean Demolisher", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Empyrean Demolisher Haste Aura",
			ActionID: core.ActionID{SpellID: 21165},
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.2)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.2)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228349/eskhandars-left-claw
	// Chance on hit: Slows enemy's movement by 60% and causes them to bleed for 150 damage over 30 sec.
	// TODO: Proc rate untested
	itemhelpers.CreateWeaponProcSpell(EskhandarsLeftClaw, "Eskhandar's Left Claw", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 22639},
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Eskhandar's Rake",
				},
				TickLength:    time.Second * 3,
				NumberOfTicks: 10,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.SpellMetrics[result.Target.UnitIndex].Hits--
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=228350/eskhandars-right-claw
	// Chance on hit: Increases your attack speed by 30% for 5 sec.
	itemhelpers.CreateWeaponProcAura(EskhandarsRightClaw, "Eskhandar's Right Claw", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Eskhandar's Rage",
			ActionID: core.ActionID{SpellID: 22640},
			Duration: time.Second * 5,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.3)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.3)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13218/fang-of-the-crystal-spider
	// Chance on hit: Slows target enemy's casting speed and increases the time between melee and ranged attacks by 10% for 10 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(FangOfTheCrystalSpider, func(agent core.Agent) {
		character := agent.GetCharacter()

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17331},
				Label:    "Fang of the Crystal Spider",
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.10)
			return aura
		})

		procMask := character.GetProcMaskForItem(FangOfTheCrystalSpider)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Fang of the Crystal Spider Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: procMask,
			PPM:      1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12590/felstriker
	// Chance on hit: All attacks are guaranteed to land and will be critical strikes for the next 3 sec.
	core.NewItemEffect(Felstriker, func(agent core.Agent) {
		character := agent.GetCharacter()

		effectAura := character.NewTemporaryStatsAura("Felstriker", core.ActionID{SpellID: 16551}, stats.Stats{stats.MeleeCrit: 100 * core.CritRatingPerCritChance, stats.MeleeHit: 100 * core.MeleeHitRatingPerHitChance}, time.Second*3)
		procMask := character.GetProcMaskForItem(Felstriker)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Felstriker Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: procMask,
			PPM:      1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				effectAura.Activate(sim)
			},
		})
	})

	core.NewItemEffect(FiendishMachete, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeElemental {
			character.PseudoStats.MobTypeAttackPower += 36
		}
	})

	itemhelpers.CreateWeaponProcSpell(FieryWarAxe, "Fiery War Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18796},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fiery War Axe Fireball",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(155, 197)
				result := spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Firebreather, "Firebreather", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16413},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 70, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 3,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Fireball",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 3, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=227934/flame-wrath
	// Chance on hit: Envelops the caster with a Fire shield for 15 sec and shoots a ring of fire dealing 130 to 170 damage to all nearby enemies.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(FlameWrath, "Flame Wrath", 1.0, func(character *core.Character) *core.Spell {
		shieldActionID := core.ActionID{SpellID: 461152}
		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         shieldActionID,
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 1, // Only the shield portion has scaling
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 10, spell.OutcomeAlwaysHit)
			},
		})
		shieldAura := character.RegisterAura(core.Aura{
			ActionID: shieldActionID,
			Label:    "Flame Wrath",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, 30)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, -30)
			},
			OnSpellHitTaken: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() {
					shieldSpell.Cast(sim, spell.Unit)
				}
			},
		})
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461151},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				shieldAura.Activate(sim)

				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(130, 170), spell.OutcomeMagicHit)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(FlurryAxe, "Flurry Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18797},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttack(sim, 1, core.ActionID{SpellID: 18797})
			},
		})
	})

	// https://www.wowhead.com/classic/item=14024/frightalon
	// Chance on hit: Lowers all attributes of target by 10 for 1 min.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frightalon, func(agent core.Agent) {
		character := agent.GetCharacter()
		procMask := character.GetProcMaskForItem(Frightalon)

		debuffAuraArray := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 19755},
				Label:    "Frightalon",
				Duration: time.Minute * 1,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -10,
						stats.Intellect: -10,
						stats.Stamina:   -10,
						stats.Spirit:    -10,
						stats.Strength:  -10,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   10,
						stats.Intellect: 10,
						stats.Stamina:   10,
						stats.Spirit:    10,
						stats.Strength:  10,
					})
				},
			})
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Frightalon Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: procMask,
			PPM:      1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuraArray.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=227994/frightskull-shaft
	// Chance on hit: Deals 8 Shadow damage every 2 sec for 30 sec and lowers their Strength for the duration of the disease.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(FrightskullShaft, "Frightskull Shaft", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18633},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot | core.SpellFlagDisease,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Weakening Disease",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, -50)
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, 50)
					},
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228029/gravestone-war-axe
	// Chance on hit: Diseases target enemy for 55 Nature damage every 3 sec for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(GravestoneWarAxe, "Gravestone War Axe", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18289},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagDisease | core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Creeping Mold",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 55, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(GryphonRidersStormhammer, "Gryphon Rider's Stormhammer", 1.0, 18081, core.SpellSchoolNature, 91, 34, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228267/gutgore-ripper
	// Chance on hit: Sends a shadowy bolt at the enemy causing 150 Shadow damage and lowering all stats by 25 for 30 sec.
	itemhelpers.CreateWeaponProcSpell(GutgoreRipper, "Gutgore Ripper", 1.0, func(character *core.Character) *core.Spell {
		procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 461682},
				Label:    "Gutgore Ripper",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -25,
						stats.Intellect: -25,
						stats.Stamina:   -25,
						stats.Spirit:    -25,
						stats.Strength:  -25,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   25,
						stats.Intellect: 25,
						stats.Stamina:   25,
						stats.Spirit:    25,
						stats.Strength:  25,
					})
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461682},
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 150, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					procAuras.Get(target).Activate(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Gutwrencher, "Gutwrencher", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16406},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Gutwrencher)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(GutRipper, "Gut Ripper", 1.0, 18107, core.SpellSchoolPhysical, 95, 26, 0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponProcDamage(HammerOfTheNorthernWind, "Hammer of the Northern Wind", 3.5, 13439, core.SpellSchoolFrost, 20, 10, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=2243/hand-of-edward-the-odd
	// Chance on hit: Next spell cast within 4 sec will cast instantly.
	itemhelpers.CreateWeaponProcAura(HandOfEdwardTheOdd, "Hand of Edward the Odd", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 18803},
			Label:    "Focus (Hand of Edward the Odd)",
			Duration: time.Second * 4,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(100000)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(1 / 100000.0)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				aura.Deactivate(sim)
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(HanzoSword, "Hanzo Sword", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=228022/headmasters-charge#comments
	// Use: Gives 20 additional intellect to party members within 30 yards. (10 Min Cooldown)
	// Originally did not stack with Arcane Intellect, but is reported to stack in SoD
	core.NewItemEffect(HeadmastersCharge, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 18264}

		buffAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Headmaster's Charge",
			Duration: time.Minute * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, 25)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, -25)
			},
		})
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 10,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	itemhelpers.CreateWeaponProcSpell(HookfangShanker, "Hookfang Shanker", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13526},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Corrosive Poison",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: -50})
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: 50})
					},
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 7, dot.OutcomeTick)
				},
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Ironfoe, "Ironfoe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 15494},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttack(sim, 2, core.ActionID{SpellID: 15494})
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(JoonhosMercy, "Joonho's Mercy", 1.0, 20883, core.SpellSchoolArcane, 70, 0, 0, core.DefenseTypeMagic)

	itemhelpers.CreateWeaponProcDamage(LinkensSwordOfMastery, "Linken's Sword of Mastery", 1.0, 18089, core.SpellSchoolNature, 45, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=227940/lord-generals-sword
	// Chance on hit: Increases attack power by 50 for 30 sec.
	// // TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(LordGeneralsSword, "Lord General's Sword", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 15602},
			Label:    "Lord General's Sword",
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       50,
					stats.RangedAttackPower: 50,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       -50,
					stats.RangedAttackPower: -50,
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=12794/masterwork-stormhammer
	// Chance on hit: Blasts up to 3 targets for 105 to 145 Nature damage.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcDamage(MasterworkStormhammer, "Masterwork Stormhammer", 0.5, 463946, core.SpellSchoolNature, 105, 40, 0.1, core.DefenseTypeMagic)

	itemhelpers.CreateWeaponProcDamage(Nightblade, "Nightblade", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19169/nightfall
	// Removed from SoD
	// core.NewItemEffect(Nightfall, func(agent core.Agent) {
	// 	makeNightfallProc(agent.GetCharacter(), "Nightfall")
	// })

	itemhelpers.CreateWeaponProcDamage(PendulumOfDoom, "Pendulum of Doom", 0.5, 10373, core.SpellSchoolPhysical, 250, 100, 0, core.DefenseTypeMelee)

	core.NewItemEffect(PipsSkinner, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeBeast {
			character.PseudoStats.MobTypeAttackPower += 45
		}
	})

	// https://www.wowhead.com/classic/item=228296/perditions-blade
	// Chance on hit: Blasts a target for 98 to 122 Fire damage.
	itemhelpers.CreateWeaponProcDamage(PerditionsBlade, "Perdition's Blade", 2.8, 461695, core.SpellSchoolFire, 98, 24, 0, core.DefenseTypeMagic)
	itemhelpers.CreateWeaponProcDamage(PerditionsBladeMolten, "Perdition's Blade", 2.8, 461695, core.SpellSchoolFire, 98, 24, 0, core.DefenseTypeMagic)

	itemhelpers.CreateWeaponProcAura(Ravager, "Ravager", 1.0, func(character *core.Character) *core.Aura {
		tickActionID := core.ActionID{SpellID: 9633}
		procActionID := core.ActionID{SpellID: 9632}
		auraActionID := core.ActionID{SpellID: 433801}

		ravegerBladestormTickSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    tickActionID,
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			BonusCoefficient: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				damage := 5.0 + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower())
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, damage, spell.OutcomeMeleeSpecialHitAndCrit)
				}
			},
		})

		character.GetOrRegisterSpell(core.SpellConfig{
			SpellSchool: core.SpellSchoolPhysical,
			ActionID:    procActionID,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagChanneled,
			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Ravager Whirlwind",
				},
				NumberOfTicks:       3,
				TickLength:          time.Second * 3,
				AffectedByCastSpeed: false,
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					ravegerBladestormTickSpell.Cast(sim, target)
				},
			},
		})

		return character.GetOrRegisterAura(core.Aura{
			Label:    "Ravager Bladestorm",
			ActionID: auraActionID,
			Duration: time.Second * 9,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.CancelAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Apply(sim)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.EnableAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Cancel(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228125/refined-arcanite-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// Chance on hit: Increases damage done by 20 and attack speed by 5% for 15 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(RefinedArcaniteChampion, func(agent core.Agent) {
		character := agent.GetCharacter()

		strengthAura := strengthOfTheChampionAura(character)
		procMask := character.GetProcMaskForItem(RefinedArcaniteChampion)
		enrageAura := enrageAura446327(character)

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Refined Arcanite Champion (Strength)",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: procMask,
			PPM:      1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				strengthAura.Activate(sim)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Refined Arcanite Champion (Enrage)",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: procMask,
			PPM:      1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				enrageAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228543/runeblade-of-baron-rivendare
	// Equip: Increases movement speed and life regeneration rate.
	// TODO: Movement speed not implemented
	core.NewItemEffect(RunebladeOfBaronRivendare, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 17625}
		healthMetrics := character.NewHealthMetrics(actionID)
		character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Unholy Aura",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 5,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.GainHealth(sim, 20, healthMetrics)
					},
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=227843/reaving-nightfall
	// Removed from SoD
	// core.NewItemEffect(ReavingNightfall, func(agent core.Agent) {
	// 	makeNightfallProc(agent.GetCharacter(), "Reaving Nightfall")
	// })

	itemhelpers.CreateWeaponProcDamage(SatyrsLash, "Satyr's Lash", 1.0, 18205, core.SpellSchoolShadow, 55, 30, 0, core.DefenseTypeMagic)

	// TODO Searing Needle adds an "Apply Aura: Mod Damage Done (Fire): 10" aura to the /target/, buffing it; not currently modelled
	itemhelpers.CreateWeaponProcDamage(SearingNeedle, "Searing Needle", 1.0, 16454, core.SpellSchoolFire, 60, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228666/seeping-willow
	// Chance on hit: Lowers all stats by 20 and deals 20 Nature damage every 3 sec to all enemies within an 8 yard radius of the caster for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(SeepingWillow, "Seeping Willow", 0.5, func(character *core.Character) *core.Spell {
		stats := stats.Stats{
			stats.Agility:   20,
			stats.Intellect: 20,
			stats.Stamina:   20,
			stats.Spirit:    20,
			stats.Strength:  20,
		}
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			return unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17196},
				Label:    "Seeping Willow",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats.Multiply(-1))
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats)
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17196},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Seeping Willow Poison",
				},
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
					debuffAuras.Get(target).Activate(sim)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
					if result.Landed() {
						spell.SpellMetrics[result.Target.UnitIndex].Hits--
						spell.Dot(aoeTarget).Apply(sim)
					}
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(SerpentSlicer, "Serpent Slicer", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17511},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Serpent Slicer)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(Shadowblade, "Shadowblade", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228272/shadowstrike
	// Chance on hit: Steals 180 to 220 life from target enemy.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(Shadowstrike, "Shadowstrike", 2.2, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 461683}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(180, 220), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(ShortswordOfVengeance, "Shortsword of Vengeance", 1.0, 13519, core.SpellSchoolHoly, 30, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228542/skullforge-reaver
	// Equip: Drains target for 2 Shadow damage every 1 sec and transfers it to the caster. Lasts for 30 sec.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(SkullforgeReaver, "Skullforge Reaver", 1.7, func(character *core.Character) *core.Spell {
		procMask := character.GetProcMaskForItem(SkullforgeReaver)
		actionID := core.ActionID{SpellID: 17484}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    procMask,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				NumberOfTicks: 30,
				TickLength:    time.Second,
				Aura: core.Aura{
					Label: "Skullforge Brand",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 2, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
					character.GainHealth(sim, result.Damage, healthMetrics)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=227886/skyriders-masterwork-stormhammer
	// Chance on hit: Blasts up to 3 targets for 105 to 145 Nature damage.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcDamage(SkyridersMasterworkStormhammer, "Skyrider's Masterwork Stormhammer", 0.5, 463946, core.SpellSchoolNature, 105, 40, 0.1, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=227683/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and purges the target's soul, increasing Fire and Holy damage taken by up to 30 and dealing an additional 75 damage over 10 sec.
	// Equip: 20% chance to deal 25 Fire damage to all nearby enemies when you are struck by a melee attack. (Proc chance: 20%)
	core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 460335},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			BonusCoefficient: .025,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 25, spell.OutcomeAlwaysHit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Immolation (Hand of Ragnaros)",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: .20,
			Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					immolationSpell.Cast(sim, aoeTarget)
				}
			},
		})

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			return unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 460338},
				Label:    "Purged by Fire",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] += 30
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] += 30
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] -= 30
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] -= 30
				},
			})
		})

		purgedByFireSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 460338},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Purged By Fire",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 5,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
					debuffAuras.Get(target).Activate(sim)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Purged by Fire Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				purgedByFireSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17182/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and an additional 75 damage over 10 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	// core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
	// 	character := agent.GetCharacter()

	// 	immolationActionID := core.ActionID{SpellID: 21142}

	// 	immolationSpell := character.GetOrRegisterSpell(core.SpellConfig{
	// 		ActionID:    immolationActionID,
	// 		SpellSchool: core.SpellSchoolFire,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		DamageMultiplier: 1,
	// 		ThreatMultiplier: 1,

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeMagicHit)
	// 		},
	// 	})

	// 	character.GetOrRegisterAura(core.Aura{
	// 		ActionID: immolationActionID,
	// 		Label:    "Immolation (Hand of Ragnaros)",
	// 		Duration: core.NeverExpires,
	// 		OnReset: func(aura *core.Aura, sim *core.Simulation) {
	// 			aura.Activate(sim)
	// 		},
	// 		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 			if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
	// 				immolationSpell.Cast(sim, spell.Unit)
	// 			}
	// 		},
	// 	})

	// 	fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
	// 		ActionID:    core.ActionID{SpellID: 21162},
	// 		SpellSchool: core.SpellSchoolFire,
	// 		DefenseType: core.DefenseTypeMagic,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		DamageMultiplier: 1,
	// 		ThreatMultiplier: 1,

	// 		Dot: core.DotConfig{
	// 			Aura: core.Aura{
	// 				Label: "Fireball (Hand of Ragnaros)",
	// 			},
	// 			TickLength:    2 * time.Second,
	// 			NumberOfTicks: 5,

	// 			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
	// 				dot.Snapshot(target, 15, isRollover)
	// 			},

	// 			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
	// 				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
	// 			},
	// 		},

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
	// 			if result.Landed() {
	// 				spell.Dot(target).Apply(sim)
	// 			}
	// 		},
	// 	})

	// 	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
	// 		Name:     "Hand of Ragnaros Trigger",
	// 		Callback: core.CallbackOnSpellHitDealt,
	// 		Outcome:  core.OutcomeLanded,
	// 		ProcMask: core.ProcMaskMelee,
	// 		PPM:      1, // Estimated based on data from WoW Armaments Discord
	// 		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 			fireballSpell.Cast(sim, result.Target)
	// 		},
	// 	})
	// })

	// https://www.wowhead.com/classic/item=227684/sulfuron-hammer
	// Chance on hit: Hurls a fiery ball that causes 83 to 101 Fire damage and an additional 16 damage over 8 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	core.NewItemEffect(SulfuronHammer, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21142},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeAlwaysHit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Immolation (Hand of Ragnaros)",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			Handler: func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
				immolationSpell.Cast(sim, spell.Unit)
			},
		})

		fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21159},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Sulfuron Hammer)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 4,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 4, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(83, 101), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Sulfuron Hammer Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // TODO: Armaments Discord didn't have any data on Sulfuron Hammer
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=227832/tempered-black-amnesty
	// Chance on hit: Reduce your threat to the current target making them less likely to attack you.
	// TODO: Proc rate untested, no way to reduce threat right now
	// itemhelpers.CreateWeaponProcSpell(TemperedBlackAmnesty, "Tempered Black Amnesty", 1.0, func(character *core.Character) *core.Spell {
	// 	return character.GetOrRegisterSpell(core.SpellConfig{
	// 		ActionID:         core.ActionID{SpellID: 23604},
	// 		SpellSchool:      core.SpellSchoolPhysical,
	// 		ProcMask:         core.ProcMaskEmpty,
	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			character.threat
	// 		},
	// 	})
	// })

	itemhelpers.CreateWeaponProcSpell(TheHandOfAntusul, "The Hand of Antu'sul", 1.0, func(character *core.Character) *core.Spell {
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				Label:    "ThunderClap-Antu'sul",
				ActionID: core.ActionID{SpellID: 13532},
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.11)
			return aura
		})

		results := make([]*core.SpellResult, min(4, character.Env.GetNumTargets()))

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13532},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 7, spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
				for _, result := range results {
					spell.DealDamage(sim, result)
					if result.Landed() {
						debuffAuras.Get(result.Target).Activate(sim)
					}
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcAura(TheJackhammer, "The Jackhammer", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "The Jackhammer Haste Aura",
			ActionID: core.ActionID{SpellID: 13533},
			Duration: time.Second * 9,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.3)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.3)
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(ThrashBlade, "Thrash Blade", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 21919},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttack(sim, 1, core.ActionID{SpellID: 21919})
			},
		})
	})

	// Not yet available
	// core.NewItemEffect(Thunderfury, func(agent core.Agent) {
	// 	character := agent.GetCharacter()

	// 	procMask := character.GetProcMaskForItem(Thunderfury)
	// 	ppmm := character.AutoAttacks.NewPPMManager(6.0, procMask)

	// 	procActionID := core.ActionID{SpellID: 21992}

	// 	singleTargetSpell := character.GetOrRegisterSpell(core.SpellConfig{
	// 		ActionID:    procActionID.WithTag(1),
	// 		SpellSchool: core.SpellSchoolNature,
	// 		DefenseType: core.DefenseTypeMagic,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		DamageMultiplier: 1,
	// 		ThreatMultiplier: 0.5,

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			spell.CalcAndDealDamage(sim, target, 300, spell.OutcomeMagicHitAndCrit)
	// 		},
	// 	})

	// 	debuffAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
	// 		return target.GetOrRegisterAura(core.Aura{
	// 			Label:    "Thunderfury",
	// 			ActionID: procActionID,
	// 			Duration: time.Second * 12,
	// 			OnGain: func(aura *core.Aura, sim *core.Simulation) {
	// 				target.AddStatDynamic(sim, stats.NatureResistance, -25)
	// 			},
	// 			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
	// 				target.AddStatDynamic(sim, stats.NatureResistance, 25)
	// 			},
	// 		})
	// 	})

	// 	results := make([]*core.SpellResult, min(5, character.Env.GetNumTargets()))

	// 	bounceSpell := character.GetOrRegisterSpell(core.SpellConfig{
	// 		ActionID:    procActionID.WithTag(2),
	// 		SpellSchool: core.SpellSchoolNature,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		ThreatMultiplier: 1,
	// 		FlatThreatBonus:  63,

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			for idx := range results {
	// 				results[idx] = spell.CalcDamage(sim, target, 0, spell.OutcomeMagicHit)
	// 				target = sim.Environment.NextTargetUnit(target)
	// 			}
	// 			for _, result := range results {
	// 				if result.Landed() {
	// 					debuffAuras[result.Target.Index].Activate(sim)
	// 				}
	// 				spell.DealDamage(sim, result)
	// 			}
	// 		},
	// 	})

	// 	character.GetOrRegisterAura(core.Aura{
	// 		Label:    "Thunderfury",
	// 		Duration: core.NeverExpires,
	// 		OnReset: func(aura *core.Aura, sim *core.Simulation) {
	// 			aura.Activate(sim)
	// 		},
	// 		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 			if !result.Landed() {
	// 				return
	// 			}

	// 			if ppmm.Proc(sim, spell.ProcMask, "Thunderfury") {
	// 				singleTargetSpell.Cast(sim, result.Target)
	// 				bounceSpell.Cast(sim, result.Target)
	// 			}
	// 		},
	// 	})
	// })

	// https://www.wowhead.com/classic/item=228273/thunderstrike
	// Chance on hit: Blasts up to 3 targets for 200 to 300 Nature damage. Each target after the first takes less damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Thunderstrike, "Thunderstrike", 1.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461686},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				initialResult := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
				// Only the initial hit can be fully resisted according to a wowhead comment
				if initialResult.Landed() {
					damageMultiplier := 1.0
					for numHits := 0; numHits < 3; numHits++ {
						spell.CalcAndDealDamage(sim, target, sim.Roll(150, 250)*damageMultiplier, spell.OutcomeMagicCrit)
						numHits++
						target = character.Env.NextTargetUnit(target)
						// TODO: Couldn't find information on what the multiplier actually is
						damageMultiplier *= .65
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=228347/typhoon
	// Chance on hit: Grants an extra attack on your next swing.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Typhoon, func(agent core.Agent) {
		character := agent.GetCharacter()
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Typhoon Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				character.AutoAttacks.ExtraMHAttack(sim, 1, core.ActionID{SpellID: 461985})
			},
		})
	})

	// https://www.wowhead.com/classic/item=13183/venomspitter
	// Chance on hit: Poisons target for 7 Nature damage every 2 sec for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Venomspitter, "Venomspitter", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18203},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Poison (Venomspitter)",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 15,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 7, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.SpellMetrics[result.Target.UnitIndex].Hits--
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcDamage(VilerendSlicer, "Vilerend Slicer", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponProcDamage(ViskagTheBloodletter, "Vis'kag the Bloodletter", 0.6, 21140, core.SpellSchoolPhysical, 240, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=227941/wraith-scythe
	// Chance on hit: Steals 45 life from target enemy.
	itemhelpers.CreateWeaponProcSpell(WraithScythe, "Wraith Scythe", 1.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 16414}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 0.3,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 45, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	///////////////////////////////////////////////////////////////////////////
	//                                 Trinkets
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=227972/burst-of-knowledge
	// Use: Reduces mana cost of all spells by 100 for 10 sec. (5 Min Cooldown)
	core.NewItemEffect(BurstOfKnowledge, func(agent core.Agent) {
		character := agent.GetCharacter()

		aura := character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			Label:    "Burst of Knowledge",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.CostMultiplier -= 1
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.CostMultiplier += 1
			},
		})

		spell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			ProcMask: core.ProcMaskEmpty,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 5,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				aura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeMana,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=228678/draconic-infused-emblem
	// Use: Increases your spell damage by up to 100 and your healing by up to 190 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(DraconicInfusedEmblem, stats.Stats{stats.SpellDamage: 128, stats.HealingPower: 236}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19288/darkmoon-card-blue-dragon
	// Equip: 2% chance on successful spellcast to allow 100% of your Mana regeneration to continue while casting for 15 sec. (Proc chance: 2%)
	core.NewItemEffect(DarkmoonCardBlueDragon, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23684}

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Aura of the Blue Dragon",
			ActionID: actionID,
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.SpiritRegenRateCasting += 1
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.SpiritRegenRateCasting -= 1
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Aura of the Blue Dragon Trigger",
			Callback:   core.CallbackOnCastComplete,
			ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
			ProcChance: .02,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19287/darkmoon-card-heroism
	// Equip: Sometimes heals bearer of 120 to 180 damage when damaging an enemy in melee.
	core.NewItemEffect(DarkmoonCardHeroism, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23689}
		healthMetrics := character.NewHealthMetrics(actionID)

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.GainHealth(sim, sim.Roll(120, 180), healthMetrics)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Heroism Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19289/darkmoon-card-maelstrom
	// Equip: Chance to strike your melee target with lightning for 200 to 300 Nature damage.
	core.NewItemEffect(DarkmoonCardMaelstrom, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23687}

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(200, 300), spell.OutcomeMagicHitAndCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Lightning Strike Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228293/essence-of-the-pure-flame
	// Equip: When struck in combat inflicts 50 Fire damage to the attacker.
	core.NewItemEffect(EssenceOfThePureFlame, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 461694},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 50, spell.OutcomeAlwaysHit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Fiery Aura Proc",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee, // TODO: Unsure if this means melee attacks or all attacks
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=227990/hand-of-injustice
	// Equip: 2% chance on ranged hit to gain 1 extra attack. (Proc chance: 2%, 2s cooldown)
	core.NewItemEffect(HandOfInjustice, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.AutoSwingRanged {
			return
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 2,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Hand of Injustice",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskRanged) && icd.IsReady(sim) && sim.Proc(0.02, "HandOfInjustice") {
					icd.Use(sim)
					aura.Unit.AutoAttacks.ExtraRangedAttack(sim, 1, core.ActionID{SpellID: 461164})
				}
			},
		})
	})

	core.NewItemEffect(HandOfJustice, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 2,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Hand of Justice",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && icd.IsReady(sim) && sim.Proc(0.02, "HandOfJustice") {
					icd.Use(sim)
					aura.Unit.AutoAttacks.ExtraMHAttack(sim, 1, core.ActionID{SpellID: 15600})
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=228599/heart-of-wyrmthalak
	// Equip: Chance to bathe your melee target in flames for 120 to 180 Fire damage.
	// TODO: Proc rate assumed from a wowhead comment and needs testing
	core.NewItemEffect(HeartOfWyrmthalak, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 462385},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(120, 180), spell.OutcomeMagicHitAndCrit)
			},
		})
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Heart of Wyrmthalak Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      0.4,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				spell.Cast(sim, result.Target)
			},
		})
	})

	// Not yet in SoD
	// core.NewItemEffect(MarkOfTheChampionPhys, func(agent core.Agent) {
	// 	character := agent.GetCharacter()

	// 	if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead || character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
	// 		character.PseudoStats.MobTypeAttackPower += 150
	// 	}
	// })

	// core.NewItemEffect(MarkOfTheChampionSpell, func(agent core.Agent) {
	// 	character := agent.GetCharacter()

	// 	if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead || character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
	// 		character.PseudoStats.MobTypeSpellPower += 85
	// 	}
	// })

	core.NewItemEffect(MarkOfTheChosen, func(agent core.Agent) {
		character := agent.GetCharacter()
		statIncrease := float64(25)
		markProcChance := 0.02

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Mark of the Chosen Effect",
			ActionID: core.ActionID{SpellID: 21970},
			Duration: time.Minute,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   statIncrease,
					stats.Agility:   statIncrease,
					stats.Strength:  statIncrease,
					stats.Intellect: statIncrease,
					stats.Spirit:    statIncrease,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   -statIncrease,
					stats.Agility:   -statIncrease,
					stats.Strength:  -statIncrease,
					stats.Intellect: -statIncrease,
					stats.Spirit:    -statIncrease,
				})
			},
		})

		core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Mark of the Chosen",
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && sim.RandomFloat("Mark of the Chosen") < markProcChance {
					procAura.Activate(sim)
				}
			},
		}))
	})

	// https://www.wowhead.com/classic/item=19812/rune-of-the-dawn
	// Equip: Increases damage done to Undead by magical spells and effects by up to 48.
	core.NewItemEffect(RuneOfTheDawn, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			character.AddStat(stats.SpellDamage, 48)
		}
	})

	core.NewItemEffect(ScarabBrooch, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: ScarabBrooch}

		shieldSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 26470},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagHelpful,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Shield: core.ShieldConfig{
				Aura: core.Aura{
					Label:    "Scarab Brooch Shield",
					Duration: time.Second * 30,
				},
			},
		})

		activeAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Persistent Shield",
			Callback: core.CallbackOnHealDealt,
			Duration: time.Second * 30,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				shieldSpell.Shield(result.Target).Apply(sim, result.Damage*0.15)
			},
		})

		spell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activeAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=227967/second-wind
	// Use: Restores 30 mana every 1 sec for 10 sec. (2 Min Cooldown)
	core.NewItemEffect(SecondWind, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 15604}
		manaMetrics := character.NewManaMetrics(actionID)
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			ProcMask: core.ProcMaskEmpty,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 1,
					NumTicks: 10,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.AddMana(sim, 30, manaMetrics)
					},
				})
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=228576/smolderwebs-eye#see-also
	// Use: Poisons target for 20 Nature damage every 2 sec for 20 sec. (2 Min Cooldown)
	core.NewItemEffect(SmolderwebsEye, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17330},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Smolderweb's Eye)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTickCounted)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=13209/seal-of-the-dawn
	// Equip: +81 Attack Power when fighting Undead.
	core.NewItemEffect(SealOfTheDawn, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			character.AddStat(stats.AttackPower, 81)
			character.AddStat(stats.AttackPower, 81)
		}
	})

	// https://www.wowhead.com/classic/item=228255/talisman-of-ephemeral-power
	// Use: Increases damage and healing done by magical spells and effects by up to 184 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TalismanOfEphemeralPower, stats.Stats{stats.SpellPower: 184}, time.Second*15, time.Second*90)

	///////////////////////////////////////////////////////////////////////////
	//                                 Other
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=228354/blazefury-medallion
	// Equip: Adds 2 fire damage to your melee attacks.
	core.NewItemEffect(BlazefuryMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 7712},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskTriggerInstant,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 2, spell.OutcomeMagicCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Blazefury Medallion Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=1168/skullflame-shield
	// Equip: When struck in combat has a 3% chance of stealing 35 life from target enemy. (Proc chance: 3%)
	// Equip: When struck in combat has a 1% chance of dealing 75 to 125 Fire damage to all targets around you. (Proc chance: 1%)
	core.NewItemEffect(SkullflameShield, func(agent core.Agent) {
		character := agent.GetCharacter()

		drainLifeActionID := core.ActionID{SpellID: 18817}
		healthMetrics := character.NewHealthMetrics(drainLifeActionID)
		drainLifeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         drainLifeActionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 35, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})

		flamestrikeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18818},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(75, 125), spell.OutcomeMagicHit)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Drain Life Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.03,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				drainLifeSpell.Cast(sim, spell.Unit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Flamestrike Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.01,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				flamestrikeSpell.Cast(sim, spell.Unit)
			},
		})
	})

	core.AddEffectsToTest = true
}

// https://www.wowhead.com/classic/spell=446327/enrage
// Used by:
// - https://www.wowhead.com/classic/item=220569/blistering-ragehammer and
// - https://www.wowhead.com/classic/item=228125/refined-arcanite-champion
func enrageAura446327(character *core.Character) *core.Aura {
	return character.GetOrRegisterAura(core.Aura{
		ActionID: core.ActionID{SpellID: 446327},
		Label:    "Enrage (446327)",
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.PseudoStats.BonusDamage += 20
			character.MultiplyAttackSpeed(sim, 1.05)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.PseudoStats.BonusDamage -= 20
			character.MultiplyAttackSpeed(sim, 1/1.05)
		},
	})
}

// TODO: This is treated as a buff, NOT a debuff in-game
// We don't have the ability to remove resistances for only one agent at a time right now
func makeBonereaversEdgeEffect(character *core.Character) *core.Spell {
	actionID := core.ActionID{SpellID: 21153}
	buffAura := character.RegisterAura(core.Aura{
		ActionID:  actionID,
		Label:     "Bonereaver's Edge",
		Duration:  time.Second * 10,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
			for _, target := range sim.Encounter.TargetUnits {
				target.AddStatDynamic(sim, stats.Armor, -700*float64(oldStacks))
				target.AddStatDynamic(sim, stats.Armor, 700*float64(newStacks))
			}
		},
	})
	return character.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			buffAura.Activate(sim)
			buffAura.AddStack(sim)
		},
	})
}

// Chance on hit: Reduces an enemy's Strength by 125 and its Stamina by 50 for 2 min.
// Equip: When struck in combat has a chance of causing the attacker to flee in terror for 2 seconds. (Proc chance: 2%)
func makeDreadbladeOfTheDestructorEffect(character *core.Character) *core.Spell {
	actionID := core.ActionID{SpellID: 462178}
	procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
		return target.GetOrRegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Enfeeble (Dreadblade of the Destructor)",
			Duration: time.Minute * 2,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:  -50,
					stats.Strength: -125,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:  50,
					stats.Strength: 125,
				})
			},
		})
	})
	return character.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			procAuras.Get(target).Activate(sim)
		},
	})
}

// Chance on hit: Spell damage taken by target increased by 15% for 5 sec.
func makeNightfallProc(character *core.Character, itemName string) {
	procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
		return target.GetOrRegisterAura(core.Aura{
			Label:    fmt.Sprintf("Spell Vulnerability (%s)", itemName),
			ActionID: core.ActionID{SpellID: 23605},
			Duration: time.Second * 5,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] *= 1.15
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] /= 1.15
			},
		})
	})

	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:     fmt.Sprintf("%s Trigger", itemName),
		Callback: core.CallbackOnSpellHitDealt,
		Outcome:  core.OutcomeLanded,
		ProcMask: core.ProcMaskMelee,
		PPM:      2,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			procAuras.Get(result.Target).Activate(sim)
		},
	})
}

func strengthOfTheChampionAura(character *core.Character) *core.Aura {
	actionID := core.ActionID{SpellID: 16916}
	healthMetrics := character.NewHealthMetrics(actionID)
	return character.GetOrRegisterAura(core.Aura{
		Label:    "Strength of the Champion",
		ActionID: actionID,
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.GainHealth(sim, sim.Roll(270, 450), healthMetrics)
			character.AddStatDynamic(sim, stats.Strength, 120)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.AddStatDynamic(sim, stats.Strength, -120)
		},
	})
}
