#lang racket

   
(define(absDiff L1 L2) 
                (cond ((and (null? L1)(null? L2))'())
                           (else(cons (abs(-(car L1)(car L2)))
                                (absDiff(cdr L1)(cdr L2))))))


(absDiff '(1 3 5 6) '(3 5 2 1))
;---------------------------------------------------------------------------------------------
(define (absDiffA L1 L2)
  (cond((and (null? L1) (null? L2)) 
        '())
       ((null? L1)
        (cons (abs (+ 0 (car L2))) (absDiffA '() (cdr L2))))
       ((null? L2)
        (cons (abs (+ (car L1) 0)) (absDiffA (cdr L1) '())))
       (else
        (cons (abs(- (car L1) (car L2))) (absDiffA (cdr L1) (cdr L2))))))

  
(absDiffA '(1 3 5 6 2 5) '(3 5 2 1))
(absDiffA '(1 3 5 6) '(3 5 2 1 2 5))
;---------------------------------------------------------------------------------------------
(define (absDiffB L1 L2)
  (cond((and (null? L1) (null? L2))
        '())
       ((null? L1)null) 
       ((null? L2) null)
       (else
        (cons (abs(- (car L1) (car L2))) (absDiffB (cdr L1) (cdr L2))))))
 
(absDiffB '(1 3 5 6 2 5) '(3 5 2 1))
(absDiffB '(1 3 5 6) '(3 5 2 1 2 5))

