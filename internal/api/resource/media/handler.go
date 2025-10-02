package media

import "net/http"

type API struct{}

// List godoc
//
//  @summary        List media
//  @description    List media
//  @tags           media
//  @accept         json
//  @produce        json
//  @success        200 {array}     DTO
//  @failure        500 {object}    err.Error
//  @router         /media [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {}

// Create godoc
//
//  @summary        Create media
//  @description    Create media
//  @tags           media
//  @accept         json
//  @produce        json
//  @param          body    body    Form    true    "Media form"
//  @success        201
//  @failure        400 {object}    err.Error
//  @failure        422 {object}    err.Errors
//  @failure        500 {object}    err.Error
//  @router         /media [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

// Get godoc
//
//  @summary        Get media
//  @description    Get media
//  @tags           media
//  @accept         json
//  @produce        json
//  @param          id	path        string  true    "Media ID"
//  @success        200 {object}    DTO
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        500 {object}    err.Error
//  @router         /media/{id} [get]
func (a *API) Get(w http.ResponseWriter, r *http.Request) {}

// Update godoc
//
//  @summary        Update media
//  @description    Update media
//  @tags           media
//  @accept         json
//  @produce        json
//  @param          id      path    string  true    "Media ID"
//  @param          body    body    Form    true    "Media form"
//  @success        200
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        422 {object}    err.Errors
//  @failure        500 {object}    err.Error
//  @router         /media/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

// Delete godoc
//
//  @summary        Delete media
//  @description    Delete media
//  @tags           media
//  @accept         json
//  @produce        json
//  @param          id  path    string  true    "Media ID"
//  @success        200
//  @failure        400 {object}    err.Error
//  @failure        404
//  @failure        500 {object}    err.Error
//  @router         /media/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
