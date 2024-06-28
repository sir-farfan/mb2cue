# mb2cue
Creates a CUE file from a MusicBrainz album definition

Installing 
----------
```
go install github.com/sir-farfan/mb2cue@latest
```

Usage
-----

Since there may be multiple versions of the same album so we need to chose a release
```
mb2cue https://musicbrainz.org/release/4ab6bd77-e501-4a54-876d-3e526aa8240c

REM release id 4ab6bd77-e501-4a54-876d-3e526aa8240c
TITLE "Ωφέλιμο φορτίο"
FILE "9712d52a-4509-3d4b-a1a2-67c88c643e31" WAVE
  TRACK 01 AUDIO
    TITLE "Τα φιλιά"
    INDEX 01 00:00:00
  TRACK 02 AUDIO
    TITLE "Δική σου εικόνα"
    INDEX 01 01:50:00
```
