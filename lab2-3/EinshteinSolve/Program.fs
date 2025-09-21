type House = {
    Position: int
    Nation: string option
    Color: string option
    Drink: string option
    Smoke: string option
    Pet: string option
}

let solve () = 
    let houses = [1..5] |> List.map(fun i -> {Position = i; Nation = None; Color = None; Drink = None; Smoke = None; Pet = None})

    let applyFixedConstraints houses = 
        houses |> List.map(fun h -> 
            match h.Position with 
            | 1 -> {h with Nation = Some "Norwegian"}
            | 3 -> {h with Drink = Some "Milk"}
            | _ -> h)

    let isValid houses =
        let find pred = houses |> List.tryFind pred
        
        [
            // Англичанин - красный дом
            match find (fun h -> h.Nation = Some "English") with
            | Some h -> h.Color = Some "Red" || h.Color = None
            | None -> true
            
            // Швед - собака
            match find (fun h -> h.Nation = Some "Swedish") with
            | Some h -> h.Pet = Some "Dog" || h.Pet = None
            | None -> true
            
            // Датчанин - чай
            match find (fun h -> h.Nation = Some "Danish") with
            | Some h -> h.Drink = Some "Tea" || h.Drink = None
            | None -> true
            
            // Зеленый дом слева от белого
            match find (fun h -> h.Color = Some "Green"), find (fun h -> h.Color = Some "White") with
            | Some g, Some w -> g.Position = w.Position - 1
            | Some g, None -> g.Position < 5
            | None, Some w -> w.Position > 1
            | None, None -> true
            
            // Зеленый дом - кофе
            match find (fun h -> h.Color = Some "Green") with
            | Some h -> h.Drink = Some "Coffee" || h.Drink = None
            | None -> true
            
            // Pall Mall - птица
            match find (fun h -> h.Smoke = Some "Pall Mall") with
            | Some h -> h.Pet = Some "Bird" || h.Pet = None
            | None -> true
            
            // Желтый дом - Dunhill
            match find (fun h -> h.Color = Some "Yellow") with
            | Some h -> h.Smoke = Some "Dunhill" || h.Smoke = None
            | None -> true
            
            // Blend рядом с котом
            match find (fun h -> h.Smoke = Some "Blend"), find (fun h -> h.Pet = Some "Cat") with
            | Some b, Some c -> abs(b.Position - c.Position) = 1
            | _ -> true
            
            // Лошадь рядом с Dunhill
            match find (fun h -> h.Pet = Some "Horse"), find (fun h -> h.Smoke = Some "Dunhill") with
            | Some h, Some d -> abs(h.Position - d.Position) = 1
            | _ -> true
            
            // Blue Master - пиво
            match find (fun h -> h.Smoke = Some "Blue Master") with
            | Some h -> h.Drink = Some "Beer" || h.Drink = None
            | None -> true
            
            // Немец - Prince
            match find (fun h -> h.Nation = Some "German") with
            | Some h -> h.Smoke = Some "Prince" || h.Smoke = None
            | None -> true
            
            // Норвежец рядом с синим домом
            match find (fun h -> h.Nation = Some "Norwegian"), find (fun h -> h.Color = Some "Blue") with
            | Some n, Some b -> abs(n.Position - b.Position) = 1
            | _ -> true
            
            // Blend рядом с водой
            match find (fun h -> h.Smoke = Some "Blend"), find (fun h -> h.Drink = Some "Water") with
            | Some b, Some w -> abs(b.Position - w.Position) = 1
            | _ -> true
        ] |> List.forall id

    let rec backtrack houses =
        if houses |> List.forall (fun h -> h.Nation.IsSome && h.Color.IsSome && h.Drink.IsSome && h.Smoke.IsSome && h.Pet.IsSome) then
            if isValid houses then Some houses else None
        else
            let emptyHouse = houses |> List.find (fun h -> 
                h.Nation.IsNone || h.Color.IsNone || h.Drink.IsNone || h.Smoke.IsNone || h.Pet.IsNone)
            
            let pos = emptyHouse.Position
            
            if emptyHouse.Nation.IsNone then
                ["English"; "Danish"; "German"; "Swedish"] |> List.tryPick (fun nation ->
                    if houses |> List.exists (fun h -> h.Nation = Some nation) then None
                    else
                        let newHouses = houses |> List.map (fun h -> 
                            if h.Position = pos then {h with Nation = Some nation} else h)
                        if isValid newHouses then backtrack newHouses else None)
            elif emptyHouse.Color.IsNone then
                ["Red"; "Green"; "White"; "Yellow"; "Blue"] |> List.tryPick (fun color ->
                    if houses |> List.exists (fun h -> h.Color = Some color) then None
                    else
                        let newHouses = houses |> List.map (fun h -> 
                            if h.Position = pos then {h with Color = Some color} else h)
                        if isValid newHouses then backtrack newHouses else None)
            elif emptyHouse.Drink.IsNone then
                ["Tea"; "Coffee"; "Beer"; "Water"] |> List.tryPick (fun drink ->
                    if houses |> List.exists (fun h -> h.Drink = Some drink) then None
                    else
                        let newHouses = houses |> List.map (fun h -> 
                            if h.Position = pos then {h with Drink = Some drink} else h)
                        if isValid newHouses then backtrack newHouses else None)
            elif emptyHouse.Smoke.IsNone then
                ["Pall Mall"; "Dunhill"; "Blend"; "Blue Master"; "Prince"] |> List.tryPick (fun smoke ->
                    if houses |> List.exists (fun h -> h.Smoke = Some smoke) then None
                    else
                        let newHouses = houses |> List.map (fun h -> 
                            if h.Position = pos then {h with Smoke = Some smoke} else h)
                        if isValid newHouses then backtrack newHouses else None)
            elif emptyHouse.Pet.IsNone then
                ["Dog"; "Bird"; "Cat"; "Horse"; "Fish"] |> List.tryPick (fun pet ->
                    if houses |> List.exists (fun h -> h.Pet = Some pet) then None
                    else
                        let newHouses = houses |> List.map (fun h -> 
                            if h.Position = pos then {h with Pet = Some pet} else h)
                        if isValid newHouses then backtrack newHouses else None)
            else None

    let initial = applyFixedConstraints houses
    match backtrack initial with
    | Some result -> 
        result |> List.iter (fun h -> 
            printfn "Дом %d: %s %s %s %s %s" 
                h.Position 
                (h.Nation |> Option.defaultValue "?")
                (h.Color |> Option.defaultValue "?")
                (h.Drink |> Option.defaultValue "?")
                (h.Smoke |> Option.defaultValue "?")
                (h.Pet |> Option.defaultValue "?"))
        let fishOwner = result |> List.find (fun h -> h.Pet = Some "Fish")
        printfn "Рыбку держит: %s" (fishOwner.Nation |> Option.defaultValue "?")
    | None -> printfn "Нет решения"

solve()
