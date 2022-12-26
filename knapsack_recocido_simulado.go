package main
import (
     "fmt" 
     "math" 
     "math/rand" 
     "time"
       )
//Creamos una función la cual generará una solucion_aleatoria inicial
func solucion_aleatoria (d int)[]int{
     rand.Seed(time.Now().UnixNano())
     var h[] int 
     i:=0
     for i <d{
	     h = append (h,rand.Intn(1+1))
	     i++
	 }
     return (h)
}
//Creamos la función para evaluar nuestras soluciones
func calcular_peso_y_valor(sol []int, pesos []int, valores []int)(int,int){
    peso:=0
    valor:=0
    for i:=0;i<len(sol);i++{
        if(sol[i]==1){
            peso=peso+pesos[i]
            valor=valor+valores[i]
        }
    }
    return peso,valor
}
//Creamos la función con la que definiremos el vecindario
func solucion_vecina(sol []int)[]int{
    Q:=sol
    rand.Seed(time.Now().UnixNano())
    min := 0
    max := len(sol)-1
    h:=rand.Intn(max - min + 1) + min
    if(Q[h]==1){
        Q[h]=0
    }else{
        Q[h]=1
    }
    return Q
}
//Creamos la función del recocido
func knapsack_recocido(capacidad int,pesos []int,valores []int, Temp_Inicial float64, Temp_Final float64, enfriamiento float64, iteraciones int) ([]int,int){
    mejor_valor := 0
    var sol []int
    solucion_actual:= solucion_aleatoria(len(pesos))
    Temp:=Temp_Inicial
    mejor_sol:=make([]int,len(pesos))
    for Temp>Temp_Final{
        i:=0
        for i< iteraciones{
            sol=solucion_vecina(solucion_actual)
            peso,valor:=calcular_peso_y_valor(sol,pesos,valores)
            if peso<=capacidad && valor>mejor_valor {
                copy(mejor_sol,sol)
                copy(solucion_actual,sol)
                mejor_valor=valor
            }else{
                var delta float64 = float64(valor-mejor_valor)
                r:=rand.Float64()
                if(r<math.Exp(-delta/Temp)){
                   copy(solucion_actual,sol) 
                }
            }
            i++
        }
        Temp=Temp*enfriamiento
    }
    return mejor_sol,mejor_valor
}
func main (){
    start0 := time.Now()
    //Guardamos en una slice los tiempos en los que encuentra cada una de las 30 soluciones
    var tiemposP01 []time.Duration
    var tiemposP07 []time.Duration
    var tiemposP08 []time.Duration
    //Guardamos las propias soluciones en una slice de slices
    var solucionesP01 [][]int
    var solucionesP07 [][]int
    var solucionesP08 [][]int
    //Guardamos los valores máximos asociados a las soluciones anteriores
    var valoresP01 []int
    var valoresP07 []int
    var valoresP08 []int
    //Evaluamos 30 veces el recocido para cada problema
    for i:=0;i<30;i++{
      peso1:= []int{23, 31, 29, 44, 53, 38, 63, 85, 89, 82}
      valor1:= []int{92, 57, 49, 68, 60, 43, 67, 84, 87, 72}
      start1 := time.Now()
      h,v:=knapsack_recocido(165,peso1,valor1,1000,0.001,0.85,100)
      elapsed1 := time.Since(start1)
      tiemposP01=append(tiemposP01,elapsed1)
      solucionesP01=append(solucionesP01,h)
      valoresP01=append(valoresP01,v)
      peso2:=[]int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
      valor2:=[]int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
      start2 := time.Now()
      j,k:=knapsack_recocido(750,peso2,valor2,1000,0.001,0.85,1000)
      elapsed2 := time.Since(start2)
      tiemposP07=append(tiemposP07,elapsed2)
      solucionesP07=append(solucionesP07,j)
      valoresP07=append(valoresP07,k)
      peso3:=[]int{382745,799601,909247,729069,467902,44328,34610,698150,823460,903959,853665,551830,610856,670702,488960,951111,323046,446298,931161,31385,496951,264724,224916,169684}
      valor3:=[]int{825594,1677009,1676628,1523970,943972,97426,69666,1296457,1679693,1902996,1844992,1049289,1252836,1319836,953277,2067538,675367,853655,1826027,65731,901489,577243,466257,369261}
      start3 := time.Now()
      l,m:=knapsack_recocido(6402560,peso3,valor3,1000,0.001,0.85,1000)
      elapsed3 := time.Since(start3)
      tiemposP08=append(tiemposP08,elapsed3)
      solucionesP08=append(solucionesP08,l)
      valoresP08=append(valoresP08,m)
    }
    //Imprimimos las respectivas matrices de cada problema
    fmt.Println("Los tiempos en los que se encontraron las soluciones de P01:")
    fmt.Println(tiemposP01)
    fmt.Println("Las soluciones de P01:")
    fmt.Println(solucionesP01)
    fmt.Println("Los valores máximos encontrados asociados a las soluciones de P01")
    fmt.Println(valoresP01)
    
    fmt.Println("Los tiempos en los que se encontraron las soluciones de P07:")
    fmt.Println(tiemposP07)
    fmt.Println("Las soluciones de P07:")
    fmt.Println(solucionesP07)
    fmt.Println("Los valores máximos encontrados asociados a las soluciones de P07")
    fmt.Println(valoresP07)

    fmt.Println("Los tiempos en los que se encontraron las soluciones de P08:")
    fmt.Println(tiemposP08)
    fmt.Println("Las soluciones de P08:")
    fmt.Println(solucionesP08)
    fmt.Println("Los valores máximos encontrados asociados a las soluciones de P08")
    fmt.Println(valoresP08)
    
    elapsed0 := time.Since(start0)
    fmt.Println("El tiempo de ejecución total de este programa: ",elapsed0 )
}

