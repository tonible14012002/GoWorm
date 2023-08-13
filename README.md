# WORM GAME

## TASKs

- [ ] Create Color Map (1 sp)
- [ ] Entity - terrain collision - (5sp)
  - [ ] Update entity physic (1 sp)
  - [ ] Handle bouncing (2 sp)
- [ ] Create Camera controller (3 sp)
  - Allow moving camera around map
- [ ] Add Player Entity - (2 sp)
  - [ ] Handle Collision with map
  - [x] Controll crosshair
  - [ ] Fire Missle
  - [ ] Toggle display players name - button M
- [ ] Create PlayerTeamManager
  - [ ] Allow get the next Player in turn
  - [ ] Get team health bar
- [ ] Fire missle ( 4 sp)

  - [ ] Create Missle entity (1 sp)
  - [ ] Allow user fire missle ( init missle velo, pos) (1 sp)
  - [ ] Handle Boomb (2 sp)
    - [ ] Remove terrain bit on bomb area
    - [ ] minus healthbar of player in boomb area

- [ ] Create `Game state machine` for handling game logic `use switch-case`
  - state 1:
    - Get next Player turn
    - Set camera track to player
    - Allow Moving camera to other area (When holding button shift, reset camera focus on player when release button)
    - `Fire Missle` -> State 2
  - State 2:
    - Set camera track to missle
    - `missle Explose || out of map` -> State 3:
  - State 3:
    - Stop camera wait 3 secs
    - Get Next Team, get next Player
    - All Player death -> State 5
  - State 5:
    - Display GameOver, Winner Team name
    - Show Menu Option `Play again | Quit`
- [ ] Implement bouncing Camera
