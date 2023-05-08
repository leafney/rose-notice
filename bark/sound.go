/**
 * @Author:      leafney
 * @Date:        2022-10-10 13:36
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

type Sound string

func (s Sound) String() string {
	return string(s)
}

const (
	SoundAlarm              Sound = "alarm.caf"
	SoundAnticipate         Sound = "anticipate.caf"
	SoundBell               Sound = "bell.caf"
	SoundBirdsong           Sound = "birdsong.caf"
	SoundBloom              Sound = "bloom.caf"
	SoundCalypso            Sound = "calypso.caf"
	SoundChime              Sound = "chime.caf"
	SoundChoo               Sound = "choo.caf"
	SoundDescent            Sound = "descent.caf"
	SoundElectronic               = "electronic.caf"
	SoundFanfare                  = "fanfare.caf"
	SoundGlass                    = "glass.caf"
	SoundGotosleep                = "gotosleep.caf"
	SoundHealthnotification       = "healthnotification.caf"
	SoundHorn                     = "horn.caf"
	SoundLadder                   = "ladder.caf"
	SoundMailsent                 = "mailsent.caf"
	SoundMinuet                   = "minuet.caf"
	SoundMultiwayinvitation       = "multiwayinvitation.caf"
	SoundNewmail                  = "newmail.caf"
	SoundNewsflash                = "newsflash.caf"
	SoundNoir                     = "noir.caf"
	SoundPaymentsuccess           = "paymentsuccess.caf"
	SoundShake                    = "shake.caf"
	SoundSherwoodforest           = "sherwoodforest.caf"
	SoundSilence                  = "silence.caf"
	SoundSpell                    = "spell.caf"
	SoundSuspense                 = "suspense.caf"
	SoundTelegraph                = "telegraph.caf"
	SoundTiptoes                  = "tiptoes.caf"
	SoundTypewriters              = "typewriters.caf"
	SoundUpdate                   = "update.caf"
)
