package main
import ("fmt" 
"math" 
"math/rand" 
"time")

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


func knapsack_recocido(capacidad int,pesos []int,valores []int, Temp_Inicial float64, Temp_Final float64, enfriamiento float64, iteraciones int) ([]int,int){
    mejor_valor := 0
    solucion_actual:= solucion_aleatoria(len(pesos))
    Temp:=Temp_Inicial
    mejor_sol:=solucion_actual
    for Temp>Temp_Final{
        i:=0
        for i< iteraciones{
            sol:=solucion_vecina(solucion_actual)
            peso,valor:=calcular_peso_y_valor(sol,pesos,valores)
            if peso<=capacidad && valor>mejor_valor {
                solucion_actual=sol
                mejor_sol=sol
                mejor_valor=valor
            }else{
                var delta float64 = float64(valor-mejor_valor)
                r:=rand.Float64()
                if(r<math.Exp(-delta/Temp)){
                   solucion_actual=sol 
                }
            }
            i++
        }
        Temp=Temp*enfriamiento
    }
    return mejor_sol,mejor_valor
}

func main (){
    peso1:= []int{23, 31, 29, 44, 53, 38, 63, 85, 89, 82}
    valor1:= []int{92, 57, 49, 68, 60, 43, 67, 84, 87, 72}
    fmt.Println(knapsack_recocido(165,peso1,valor1,1000,0.001,0.85,100))
}
