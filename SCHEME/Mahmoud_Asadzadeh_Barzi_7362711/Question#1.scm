#lang racket
(define (distanceGPS Lat1 Lon1 Lat2 Lon2)
  (let ([Lat11 (* pi(/ Lat1 180))]
        [Lon11 (* pi(/ Lon1 180))]
        [Lat22 (* pi(/ Lat2 180))]
        [Lon22 (* pi(/ Lon2 180))])
  (*(* 2 (sin (sqrt(+(expt(sin(/ (- Lat11 Lat22) 2)) 2)
                     (*(*(cos Lat11) 
                         (cos Lat22))
                       (expt(sin(/ (- Lon11 Lon22) 2)) 2))))))6371)))


(distanceGPS 45.421016 -75.690018 45.4222 -75.6824)