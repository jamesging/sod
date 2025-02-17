package druid

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/proto"
)

/* TODO: Bear mangle
func (druid *Druid) registerMangleBearSpell() {
	if !druid.Talents.Mangle {
		return
	}

	mangleAuras := druid.NewEnemyAuraArray(core.MangleAura)
	durReduction := (0.5) * float64(druid.Talents.ImprovedMangle)

	druid.MangleBear = druid.RegisterSpell(Bear, core.SpellConfig{
		SpellCode:   SpellCode_DruidMangleBear
		ActionID:    core.ActionID{SpellID: 48564},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       SpellFlagOmen | core.SpellFlagMeleeMetrics | core.SpellFlagIncludeTargetBonusDamage | core.SpellFlagAPL,

		RageCost: core.RageCostOptions{
			Cost:   20 - float64(druid.Talents.Ferocity),
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Duration(float64(time.Second) * (6 - durReduction)),
			},
		},

		DamageMultiplier: (1 + 0.1*float64(druid.Talents.SavageFury)) * 1.15
		ThreatMultiplier: core.TernaryFloat64(druid.HasSetBonus(ItemSetThunderheartHarness, 2), 1.15, 1),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 299/1.15 +
				spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower()) +
				spell.BonusWeaponDamage()

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if result.Landed() {
				mangleAuras.Get(target).Activate(sim)
			} else {
				spell.IssueRefund(sim)
			}

			if druid.BerserkAura.IsActive() {
				spell.CD.Reset()
			}
		},

		RelatedAuras: []core.AuraArray{mangleAuras},
	})
        }
*/

func (druid *Druid) registerMangleCatSpell() {
	if !druid.HasRune(proto.DruidRune_RuneHandsMangle) {
		return
	}

	hasGoreRune := druid.HasRune(proto.DruidRune_RuneHelmGore)

	weaponMulti := 2.7
	energyCost := 40 - float64(druid.Talents.Ferocity)

	mangleAuras := druid.NewEnemyAuraArray(core.MangleAura)
	druid.MangleCat = druid.RegisterSpell(Cat, core.SpellConfig{
		SpellCode:   SpellCode_DruidMangleCat,
		ActionID:    core.ActionID{SpellID: 409828},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       SpellFlagOmen | core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		EnergyCost: core.EnergyCostOptions{
			Cost:   energyCost,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
		},

		DamageMultiplier: (1 + 0.1*float64(druid.Talents.SavageFury)) * weaponMulti,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower())

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if result.Landed() {
				druid.AddComboPoints(sim, 1, spell.ComboPointMetrics())
				mangleAuras.Get(target).Activate(sim)

				if hasGoreRune {
					druid.rollGoreCatReset(sim)
				}
			} else {
				spell.IssueRefund(sim)
			}
		},

		RelatedAuras: []core.AuraArray{mangleAuras},
	})
}

func (druid *Druid) CurrentMangleCatCost() float64 {
	return druid.MangleCat.ApplyCostModifiers(druid.MangleCat.DefaultCast.Cost)
}

func (druid *Druid) IsMangle(spell *core.Spell) bool {
	if druid.MangleBear != nil && druid.MangleBear.IsEqual(spell) {
		return true
	} else if druid.MangleCat != nil && druid.MangleCat.IsEqual(spell) {
		return true
	}
	return false
}
