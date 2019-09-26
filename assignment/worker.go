package main

import (
	"fmt"
)

type Worker struct { //Worker body contains ID,work,workqueue,quit signal
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker { //creating workers with different ID

	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

func (w *Worker) Start() { //Starts a worker if work is present in its queue else quits
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case newentry := <-w.Work:

				var lastinsertedid int64
				//inserting into database and returning last inserted id
				sqlStatement := `INSERT INTO locale1(
						user_id,
						vehicle_model_id ,
						package_id ,
						travel_type_id ,
						from_area_id ,
						to_area_id ,
						from_city_id ,
						to_city_id ,
						from_date ,
						to_date ,
						online_booking ,
						mobile_site_booking ,
						booking_created ,
						from_lat ,
						from_long ,
						to_lat ,
						to_long ,
						car_cancellation )
						VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)
						RETURNING id`

				err := db.QueryRow(sqlStatement,
					newentry.Userid,
					newentry.Vehiclemodelid,
					newentry.Packageid,
					newentry.Traveltypeid,
					newentry.Fromareaid,
					newentry.Toareaid,
					newentry.Fromcityid,
					newentry.Tocityid,
					newentry.Fromdate,
					newentry.Todate,
					newentry.Onlinebooking,
					newentry.Mobilesitebooking,
					newentry.Bookingcreated,
					newentry.Fromlat,
					newentry.Fromlong,
					newentry.Tolat,
					newentry.Tolong,
					newentry.Carcancellation).Scan(&lastinsertedid)

				if err != nil {
					fmt.Println(err.Error())
				}
				// inserting last inserted id to cache since cache  by concurrent requests locks are required for consistency
				shared.mu.Lock()
				_, ok := shared.cache[lastinsertedid]

				if ok != true {
					shared.cache[lastinsertedid] = true
				}
				shared.mu.Unlock()

				//returns connection to pool
				defer db.Close()

				fmt.Printf("Inserted with ID=%d\n", lastinsertedid)

			case <-w.QuitChan:
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
