package main

type WorkRequest struct { //request body
	ID                int64   `json:"Id"`
	Userid            string  `json:"Userid" validate:"required"`
	Vehiclemodelid    int     `json:"Vehiclemodelid" validate:"required,numeric"`
	Packageid         int     `json:"Packageid" validate:"numeric"`
	Traveltypeid      int     `json:"Traveltypeid" validate:"required,numeric"`
	Fromareaid        int     `json:"Fromareaid" validate:"required,numeric"`
	Toareaid          int     `json:"Toareaid" validate:"numeric"`
	Fromcityid        int     `json:"Fromcityid" validate:"numeric"`
	Tocityid          int     `json:"Tocityid" validate:"numeric"`
	Fromdate          string  `json:"Fromdate" validate:"required"`
	Todate            string  `json:"Todate"`
	Onlinebooking     int     `json:"Onlinebooking" validate:"required,numeric"`
	Mobilesitebooking int     `json:"Mobilesitebooking" validate:"required,numeric"`
	Bookingcreated    string  `json:"Bookingcreated" validate:"required"`
	Fromlat           float32 `json:"Fromlat" validate:"required,latitude"`
	Fromlong          float32 `json:"Fromlong" validate:"required,longitude"`
	Tolat             float32 `json:"Tolat" validate:"latitude"`
	Tolong            float32 `json:"Tolong" validate:"longitude"`
	Carcancellation   int     `json:"Carcancellation" validate:"required"`
}
