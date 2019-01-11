
absDiff([],[],[]).
absDiff([H1|T1],[H2|T2],L):- TH3 is abs(H2-H1),	
							absDiff(T1,T2,L1),
							L= [TH3|L1].

%-----------------------------------------------------------------
%a)
absDiffA([],[],[]).
absDiffA([H4|T4],[], L):- TH4 is abs(H4 - 0),
							absDiffA(T4,[],L1),
							L= [TH4|L1].
							
absDiffA([],[H5|T5],L):-TH5 is abs(H5-0),
							absDiffA(T5,[],L1),
							L= [TH5|L1].

absDiffA([H1|T1],[H2|T2],L):- TH3 is H2 -H1,
							H3 is abs(TH3),	
							absDiffA(T1,T2,L1),
							L= [H3|L1].

%-----------------------------------------------------------------
%b)
absDiffB([],[],[]).
absDiffB([H4|_],[],[]):- !.
absDiffB([],[H5|_],[]):- !.
absDiffB([H1|T1],[H2|T2],L):- TH3 is H2 -H1,
							H3 is abs(TH3),	
							absDiffB(T1,T2,L1),
							L= [H3|L1].

