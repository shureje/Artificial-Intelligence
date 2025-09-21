% База знаний - возможные значения
color(red). color(green). color(white). color(blue). color(yellow).
nation(brit). nation(swede). nation(dane). nation(norwegian). nation(german).
drink(tea). drink(coffee). drink(milk). drink(beer). drink(water).
smoke(pall_mall). smoke(dunhill). smoke(marlboro). smoke(winfield). smoke(rothmans).
pet(dog). pet(bird). pet(cat). pet(horse). pet(fish).

% Ограничения из условий задачи
constraint1(Houses) :- nth1(1, Houses, house(_, norwegian, _, _, _)).  % 1. Норвежец живёт в первом доме
constraint2(Houses) :- member(house(red, brit, _, _, _), Houses).      % 2. Англичанин живёт в красном доме
constraint3(Houses) :- nextto(house(green,_,_,_,_), house(white,_,_,_,_), Houses).  % 3. Зелёный дом слева от белого
constraint4(Houses) :- member(house(_, dane, tea, _, _), Houses).       % 4. Датчанин пьёт чай
constraint5(Houses) :- 
    (nextto(house(_,_,_,marlboro,_), house(_,_,_,_,cat), Houses) ;
     nextto(house(_,_,_,_,cat), house(_,_,_,marlboro,_), Houses)).      % 5. Курильщик Marlboro рядом с владельцем кошек
constraint6(Houses) :- member(house(yellow, _, _, dunhill, _), Houses). % 6. В жёлтом доме курят Dunhill
constraint7(Houses) :- member(house(_, german, _, rothmans, _), Houses). % 7. Немец курит Rothmans
constraint8(Houses) :- nth1(3, Houses, house(_, _, milk, _, _)).        % 8. В центре пьют молоко
constraint9(Houses) :- 
    (nextto(house(_,_,_,marlboro,_), house(_,_,water,_,_), Houses) ;
     nextto(house(_,_,water,_,_), house(_,_,_,marlboro,_), Houses)).    % 9. Сосед курильщика Marlboro пьёт воду
constraint10(Houses) :- member(house(_, _, _, pall_mall, bird), Houses). % 10. Курильщик Pall Mall выращивает птиц
constraint11(Houses) :- member(house(_, swede, _, _, dog), Houses).      % 11. Швед выращивает собак
constraint12(Houses) :- 
    (nextto(house(_,norwegian,_,_,_), house(blue,_,_,_,_), Houses) ;
     nextto(house(blue,_,_,_,_), house(_,norwegian,_,_,_), Houses)).    % 12. Норвежец рядом с синим домом
constraint13(Houses) :- member(house(blue, _, _, _, horse), Houses).     % 13. В синем доме выращивают лошадей
constraint14(Houses) :- member(house(_, _, beer, winfield, _), Houses).  % 14. Курильщик Winfield пьёт пиво
constraint15(Houses) :- member(house(green, _, coffee, _, _), Houses).   % 15. В зелёном доме пьют кофе

% Главное правило решения
solve(Owner) :-
    Houses = [house(C1,N1,D1,S1,P1), house(C2,N2,D2,S2,P2), house(C3,N3,D3,S3,P3), house(C4,N4,D4,S4,P4), house(C5,N5,D5,S5,P5)],

    permutation([red,green,white,blue,yellow], [C1,C2,C3,C4,C5]),
    constraint1(Houses), constraint2(Houses), constraint3(Houses),

    permutation([brit,swede,dane,norwegian,german], [N1,N2,N3,N4,N5]),
    constraint4(Houses), constraint5(Houses), constraint6(Houses),

    permutation([tea,coffee,milk,beer,water], [D1,D2,D3,D4,D5]),
    constraint7(Houses), constraint8(Houses), constraint9(Houses),

    permutation([pall_mall,dunhill,marlboro,winfield,rothmans], [S1,S2,S3,S4,S5]),
    constraint10(Houses), constraint11(Houses), constraint12(Houses),

    permutation([dog,bird,cat,horse,fish], [P1,P2,P3,P4,P5]),
    constraint13(Houses), constraint14(Houses), constraint15(Houses),   
    member(house(_, Owner, _, _, fish), Houses).

% Вспомогательное правило для показа полного решения
show_solution :-
    Houses = [house(C1,N1,D1,S1,P1), house(C2,N2,D2,S2,P2), house(C3,N3,D3,S3,P3), house(C4,N4,D4,S4,P4), house(C5,N5,D5,S5,P5)],

   permutation([red,green,white,blue,yellow], [C1,C2,C3,C4,C5]),
    constraint1(Houses), constraint2(Houses), constraint3(Houses),

    permutation([brit,swede,dane,norwegian,german], [N1,N2,N3,N4,N5]),
    constraint4(Houses), constraint5(Houses), constraint6(Houses),

    permutation([tea,coffee,milk,beer,water], [D1,D2,D3,D4,D5]),
    constraint7(Houses), constraint8(Houses), constraint9(Houses),

    permutation([pall_mall,dunhill,marlboro,winfield,rothmans], [S1,S2,S3,S4,S5]),
    constraint10(Houses), constraint11(Houses), constraint12(Houses),

    permutation([dog,bird,cat,horse,fish], [P1,P2,P3,P4,P5]),
    constraint13(Houses), constraint14(Houses), constraint15(Houses),  
    write('Решение:'), nl,
    print_houses(Houses, 1).

print_houses([], _).
print_houses([house(Color, Nation, Drink, Smoke, Pet)|T], N) :-
    write('Дом '), write(N), write(': '),
    write(Color), write(' '), write(Nation), write(' '), 
    write(Drink), write(' '), write(Smoke), write(' '), write(Pet), nl,
    N1 is N + 1,
    print_houses(T, N1).
