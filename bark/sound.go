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
	SoundElectronic         Sound = "electronic.caf"
	SoundFanfare            Sound = "fanfare.caf"
	SoundGlass              Sound = "glass.caf"
	SoundGotosleep          Sound = "gotosleep.caf"
	SoundHealthnotification Sound = "healthnotification.caf"
	SoundHorn               Sound = "horn.caf"
	SoundLadder             Sound = "ladder.caf"
	SoundMailsent           Sound = "mailsent.caf"
	SoundMinuet             Sound = "minuet.caf"
	SoundMultiwayinvitation Sound = "multiwayinvitation.caf"
	SoundNewmail            Sound = "newmail.caf"
	SoundNewsflash          Sound = "newsflash.caf"
	SoundNoir               Sound = "noir.caf"
	SoundPaymentsuccess     Sound = "paymentsuccess.caf"
	SoundShake              Sound = "shake.caf"
	SoundSherwoodforest     Sound = "sherwoodforest.caf"
	SoundSilence            Sound = "silence.caf"
	SoundSpell              Sound = "spell.caf"
	SoundSuspense           Sound = "suspense.caf"
	SoundTelegraph          Sound = "telegraph.caf"
	SoundTiptoes            Sound = "tiptoes.caf"
	SoundTypewriters        Sound = "typewriters.caf"
	SoundUpdate             Sound = "update.caf"
)
