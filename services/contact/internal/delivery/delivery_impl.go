package delivery

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ContactDeliveryImpl struct {
	usecase usecase.ContactUseCase
}

func NewContactDelivery(usecase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{
		usecase: usecase,
	}
}

func (c *ContactDeliveryImpl) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.usecase.CreateContact(contact)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (cd *ContactDeliveryImpl) GetContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, "Contact ID should be integer value!", http.StatusBadRequest)
        return
    }

    contact, err := cd.usecase.GetContact(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(contact)
}

func (cd *ContactDeliveryImpl) UpdateContact(w http.ResponseWriter, r *http.Request) {
    var contact domain.Contact

    if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := cd.usecase.UpdateContact(contact)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(contact)
}

func (cd *ContactDeliveryImpl) DeleteContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, "Contact ID should be integer value!", http.StatusBadRequest)
        return
    }

    err = cd.usecase.DeleteContact(id)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDeliveryImpl) CreateGroup(w http.ResponseWriter, r *http.Request) {
    var group domain.Group

    if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := cd.usecase.CreateGroup(group)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (cd *ContactDeliveryImpl) GetGroup(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/groups/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }

    group, err := cd.usecase.GetGroup(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(group)
}

func (cd *ContactDeliveryImpl) AddContactToGroup(w http.ResponseWriter, r *http.Request) {
    var request struct {
        ContactID int `json:"contact_id"`
        GroupID   int `json:"group_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := cd.usecase.AddContactToGroup(request.ContactID, request.GroupID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"result": "Contact added to group successfully"})
}
