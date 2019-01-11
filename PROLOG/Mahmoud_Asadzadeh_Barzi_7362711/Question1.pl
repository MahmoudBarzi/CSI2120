%a)
findall(X,(person(X,Y),employee(X,Z),company(Z,W),\+(Y=W)),L).
%b)
findall(X,(company(X,Y),city(Y,ontario)),L).
%c)
findall(X,(person(X,_),\+employee(X,_)),L).
%d)
findall(X,(employee(X,Y),company(Y,ottawa)),L).
%e)
setof(X,Y^(employee(X,Y),company(Y,ottawa)),L).

