#lang racket

(define (children E T)
  (define (getForestCars L)
    (cond
      ((null? L)'())
      ((list?(car L))
       (list(caar L)(getForestCars(cdr L))))))
  (define (getTreeCars T)
    (cond
      ((null? T)'()) 
      ((and(number? (car T))(null? (cdr T)))'())
      (not(null? (cdr T))
          (list (caadr T)(getForestCars(cddr T))))))
  (cond
    ((null? T)'())
    ((and(number?(car T))(eq?(car T) E))
     (sort(flatten(getTreeCars T))>)) 
    ((and(number?(car T))(not(eq?(car T) E)))
     (children E (cdr T)))
    ((and(list?(car T))(eq?(caar T) E))
     (sort(flatten(getTreeCars (car T)))>))
    ((and(list?(car T))(not(eq?(caar T) E)))
     (if(not(null?(children E (car T))))
        (children E (car T))
        (children E (cdr T))))))

(define atree '(10
                (2 (4 (9 (3))
                      (12 (1 (2)))
                      (16)))
                (5 (7)
                   (21))
                (6)))
(children 5 atree)

(define anothertree '(1
                      (2 (6) (7) (8))
                      (3)
                      (4 (9 (12)))
                      (5 (10) (11))))
;(children 12 anothertree)
