package handlers

import (
	s "github.com/0B1t322/Note-API/note"
	"log"
	"strconv"
	"strings"

	"github.com/0B1t322/Note-API/models/note"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type NoteHandler struct {
	db *gorm.DB
}

func New(db *gorm.DB) s.NoteService {
	return &NoteHandler{
		db: db,
	}
}

// Add godoc
// @Summary Add
// @Description add note
// @Accept json
// @Produce json
// @Success 201 {object} note.Note
// @Failure 500 
// @Router /notes [post]
func(ns *NoteHandler) Add(w http.ResponseWriter, r *http.Request) {
	note := &note.Note{}
	d := json.NewDecoder(r.Body)

	if err := d.Decode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	if err := ns.db.Create(note).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	data, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// GetAll godoc
// @Summary GetAll
// @Description return all notes
// @Produce json
// @Success 200 {Note} {object} note.Note
// @Failure 500
// @Router /notes [get]
func(ns *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	query := vars.Get("query")

	notes := []note.Note{}

	if err := ns.db.Find(&notes).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	if query != "" {
		log.Println("query not empty ")
		notes = func(notes []note.Note ,query string) []note.Note {
			qNotes := []note.Note{}
			for _, n := range notes {
				if strings.ContainsAny(n.Content, query) || strings.ContainsAny(n.Title, query) {
					qNotes = append(qNotes, n)
				}
			}

			return qNotes
		}(notes, query)
	}

	data, err := json.Marshal(&notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	w.Write(data)
}

// Get godoc
// @Summary Get
// @Description Get note by id
// @Produce json
// @Success 200 {object} note.Note
// @Failure 404
// @Failure 500 
// @Router /notes/:id [get]
func(ns *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id := vars["id"]

	id := func() uint {
		u, _ := strconv.ParseUint(_id, 10, 64)
		return uint(u)
	}()

	note := &note.Note{}

	if err := ns.db.First(note, "id = ?", id).Error; err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	data, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	w.Write(data)
}

// Update godoc
// @Summary Update
// @Description Update note
// @Accept json
// @Success 200
// @Failure 500 
// @Router /notes/:id [put]
func(ns *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id := vars["id"]

	id := func() uint {
		u, _ := strconv.ParseUint(_id, 10, 64)
		return uint(u)
	}()

	note := &note.Note{}

	d := json.NewDecoder(r.Body)

	if err := d.Decode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}

	note.ID = id

	if err := ns.db.Updates(note).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}
}

// Delete godoc
// @Summary Delete
// @Description delete note
// @Success 200
// @Failure 404
// @Failure 500
// @Router /notes/:id [delete]
func(ns *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id := vars["id"]

	id := func() uint {
		u, _ := strconv.ParseUint(_id, 10, 64)
		return uint(u)
	}()

	if err := ns.db.Delete(&note.Note{ID: id}).Error; err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return
	}
}

