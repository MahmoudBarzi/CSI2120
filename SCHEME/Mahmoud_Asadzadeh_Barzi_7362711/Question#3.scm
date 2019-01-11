#lang racket

(define (duplicatePair lst)
  (define counter
    (lambda (E L)
      (cond ((null? L)0)
            ((eqv? E (car L))(+ 1 (counter E (cdr L)))  )
            (else(counter E (cdr L))))))
  (define deleteAll
   (lambda (E L)
      (if (null? L )
            '()
      (if (eqv? (car L) E)
           (deleteAll E (cdr L))
      (cons (car L) (deleteAll E (cdr L)))))))
  
  (if(null? lst)
      '()
      (append (list(cons (car lst) (+ 1 (counter (car lst)(cdr lst)))))
              (duplicatePair (deleteAll (car lst) lst)))))

(duplicatePair '(1 a 5 6 2 b a 5 5))
(duplicatePair '(b a a a))
;-----------------------------------------------------------------------------------------
(define (duplicateList lst)
  (define counter
    (lambda (E L)
      (cond ((null? L)0)
            ((eqv? E (car L))(+ 1 (counter E (cdr L)))  )
            (else(counter E (cdr L))))))
  (define deleteAll
   (lambda (E L)
      (if (null? L )
            '()
      (if (eqv? (car L) E)
           (deleteAll E (cdr L))
      (cons (car L) (deleteAll E (cdr L)))))))
  
  (if(null? lst)
      '()
      (append (list(list (car lst) (+ 1 (counter (car lst)(cdr lst)))))
              (duplicateList (deleteAll (car lst) lst)))))

(duplicateList '(1 a 5 6 2 b a 5 5))
(duplicateList '(b a a a))
;-----------------------------------------------------------------------------------------
(define (duplicateListSorted liste)
  (define (duplicateList lst)
    (define counter
      (lambda (E L)
        (cond ((null? L)0)
              ((eqv? E (car L))(+ 1 (counter E (cdr L)))  )
              (else(counter E (cdr L))))))
    (define deleteAll
      (lambda (E L)
        (if (null? L )
            '()
            (if (eqv? (car L) E)
                (deleteAll E (cdr L))
                (cons (car L) (deleteAll E (cdr L)))))))
  
    (if(null? lst)
       '()
       (append (list(list (car lst) (+ 1 (counter (car lst)(cdr lst)))))
               (duplicateList (deleteAll (car lst) lst)))))
  (define (sorting lst)
    (sort lst
          (lambda (L1 L2) (> (cadr L1) (cadr L2)))))
  (sorting(duplicateList liste))
  
  )
(duplicateListSorted '(1 a 5 6 2 b a 5 5))
(duplicateListSorted '(b a a a))





































