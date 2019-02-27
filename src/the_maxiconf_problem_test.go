// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package maxibonkata

import (
	"testing"
	"github.com/flyingmutant/rapid"
)

const MAX_MELTED_MAXIBONS int = 1

// theMaxiconf keeps track of the maxibons enjoyed by the attendants
// (Developer instances) at the venue (KarumiHQs instance).
type theMaxiconf struct {
	awesome_conf bool
	melted_maxibons int
	venue *KarumiHQs
	developers map[string]Developer
	grabbings map[string]bool
}

// Init is an action for initializing  TheMaxiconfProblem required data.
func (m *theMaxiconf) Init() func(*rapid.T) {
	return rapid.Bind(m.init)
}

func (m *theMaxiconf) init(*rapid.T) {
	var venue = NewKarumiHQs( "The Maxiconf Fridge" )
	var developers = MakeDevelopers()

	m.venue = &venue
	m.developers = developers
	m.grabbings =  make(map[string]bool)
	for k, _ := range developers {
		m.grabbings[k] = false
	}
}

func (mc *theMaxiconf) totalGrabbings() (t int) {
	for _, v := range mc.grabbings {
		if v {
			t += 1
		}
	}
	return
}

// Check verifies that all required invariants hold.
func (mc *theMaxiconf) Check(t *rapid.T) {
	if (mc.totalGrabbings() == len(mc.grabbings)) &&
			(mc.venue.MeltedMaxibons() <= MAX_MELTED_MAXIBONS) {
		if mc.venue.MeltedMaxibons() == 0 {
			t.Fatalf("AWESOME!! No melted maxibons after %v grabbings.",
				mc.totalGrabbings())

		} else {
			t.Fatalf("Delivered %v melted maxibons after %v grabbings.",
				mc.venue.MeltedMaxibons(),
				mc.totalGrabbings())
		}
	}
}

// Developers
func (m *theMaxiconf) Pedro_grabs_3() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["pedro"]),  m.pedro )
}
func (m *theMaxiconf) pedro(t *rapid.T) {
	m.venue.OpenFridge( m.developers["pedro"] )
	m.grabbings["pedro"] = true
}

func (m *theMaxiconf) Fran_grabs_1() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["fran"]), m.fran )
}
func (m *theMaxiconf) fran(t *rapid.T) {
	m.venue.OpenFridge( m.developers["fran"] )
	m.grabbings["fran"] = true
}

func (m *theMaxiconf) Davide_grabs_0() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["davide"]),  m.davide )
}
func (m *theMaxiconf) davide(t *rapid.T) {
	m.venue.OpenFridge( m.developers["davide"] )
	m.grabbings["davide"] = true
}

func (m *theMaxiconf) Sergio_grabs_2() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["sergio"]),  m.sergio )
}
func (m *theMaxiconf) sergio(t *rapid.T) {
	m.venue.OpenFridge( m.developers["sergio"] )
	m.grabbings["sergio"] = true
}

func (m *theMaxiconf) Jorge_grabs_1() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["jorge"]),  m.jorge )
}
func (m *theMaxiconf) jorge(t *rapid.T) {
	m.venue.OpenFridge( m.developers["jorge"] )
	m.grabbings["jorge"] = true
}

func (m *theMaxiconf) Nuria_grabs_3() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["nuria"]),  m.nuria )
}
func (m *theMaxiconf) nuria(t *rapid.T) {
	m.venue.OpenFridge( m.developers["nuria"] )
	m.grabbings["nuria"] = true
}

func (m *theMaxiconf) Fausto_grabs_2() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["fausto"]),  m.fausto )
}
func (m *theMaxiconf) fausto(t *rapid.T) {
	m.venue.OpenFridge( m.developers["fausto"] )
	m.grabbings["fausto"] = true
}

func (m *theMaxiconf) Julia_grabs_0() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["julia"]),  m.julia )
}
func (m *theMaxiconf) julia(t *rapid.T) {
	m.venue.OpenFridge( m.developers["julia"] )
	m.grabbings["julia"] = true
}

func (m *theMaxiconf) Luismi_grabs_3() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["luismi"]),  m.luismi )
}
func (m *theMaxiconf) luismi(t *rapid.T) {
	m.venue.OpenFridge( m.developers["luismi"] )
	m.grabbings["luismi"] = true
}

func (m *theMaxiconf) Susana_grabs_2() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["susana"]),  m.susana )
}
func (m *theMaxiconf) susana(t *rapid.T) {
	m.venue.OpenFridge( m.developers["susana"] )
	m.grabbings["susana"] = true
}

func (m *theMaxiconf) Sahu_grabs_1() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["sahu"]),  m.sahu )
}
func (m *theMaxiconf) sahu(t *rapid.T) {
	m.venue.OpenFridge( m.developers["sahu"] )
	m.grabbings["sahu"] = true
}

func (m *theMaxiconf) Vero_grabs_3() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["vero"]),  m.vero )
}
func (m *theMaxiconf) vero(t *rapid.T) {
	m.venue.OpenFridge( m.developers["vero"] )
	m.grabbings["vero"] = true
}

func (m *theMaxiconf) Vito_grabs_3() func(*rapid.T) {
	return rapid.BindIf( !(m.grabbings["vito"]),  m.vito )
}
func (m *theMaxiconf) vito(t *rapid.T) {
	m.venue.OpenFridge( m.developers["vito"] )
	m.grabbings["vito"] = true
}

// Go!!
func TestTheMaxiconfProblemWithRapid(t *testing.T) {
	rapid.Check(t, rapid.StateMachine(&theMaxiconf{}))
}
