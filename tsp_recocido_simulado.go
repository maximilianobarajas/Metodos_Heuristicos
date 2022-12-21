/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import ("fmt"
        "math"
        "math/rand"
        "time"
        )
func solucion_aleatoria(d int) []int{
    var h []int
    rand.Seed(time.Now().Unix())
    b:=rand.Perm(d-1)
    for i := range b {
      b[i]++
    }
    h=append(h,0)
    h=append(h,b...)
    return(h)
} 
func solucion_vecina(sol []int) []int{
    Q:=sol
    rand.Seed(time.Now().UnixNano())
	j := rand.Intn(len(sol)-1)+1
	i:= rand.Intn(len(sol)-1)+1
	Q[i], Q[j] = Q[j], Q[i]
	return Q
}  
func costo(matriz_distancias [][]int, sol []int) float64{
    var costo float64 =0
    rango:=len(sol)-1
    i:=0
    for i < rango{
        costo=costo+ float64(matriz_distancias[sol[i]][sol[i+1]])
        i++
    }
    costo=costo +float64(matriz_distancias[sol[len(sol)-1]][0])
    return costo
}
func tsp_recocido(matriz_distancias [][]int, Temp_Inicial float64, Temp_Final float64, enfriamiento float64, iteraciones int) ([]int,float64){
    mejor_costo := math.Inf(1)
    solucion_actual:= solucion_aleatoria(len(matriz_distancias))
    Temp:=Temp_Inicial
    mejor_sol:=solucion_actual
    for Temp>Temp_Final{
        i:=0
        for i< iteraciones{
            sol:=solucion_vecina(solucion_actual)
            if costo(matriz_distancias,sol)<mejor_costo {
                solucion_actual=sol
                mejor_sol=sol
                mejor_costo=costo(matriz_distancias,sol)
                
            }else{
                var delta float64 = float64(costo(matriz_distancias,sol)-costo(matriz_distancias, solucion_actual))
                r:=rand.Float64()
                if(r<math.Exp(-delta/Temp)){
                   solucion_actual=sol 
                }
            }
            i++
        }
        Temp=Temp*enfriamiento
    }
    return mejor_sol,mejor_costo
}
func main(){
  caso1 := [][]int{
   {9999, 3, 5, 48, 48, 8, 8, 5, 5, 3, 3, 0, 3, 5, 8, 8, 5},
  {3, 9999, 3, 48, 48, 8, 8, 5, 5, 0, 0, 3, 0, 3, 8, 8, 5},
  {5, 3, 9999, 72, 72, 48, 48, 24, 24, 3, 3, 5, 3, 0, 48, 48, 24},
  {48, 48, 74, 9999, 0, 6, 6, 12, 12, 48, 48, 48, 48, 74, 6, 6, 12},
  {48, 48, 74, 0, 9999, 6, 6, 12, 12, 48, 48, 48, 48, 74, 6, 6, 12},
  {8, 8, 50, 6, 6, 9999, 0, 8, 8, 8, 8, 8, 8, 50, 0, 0, 8},
  {8, 8, 50, 6, 6, 0, 9999, 8, 8, 8, 8, 8, 8, 50, 0, 0, 8},
  {5, 5, 26, 12, 12, 8, 8, 9999, 0, 5, 5, 5, 5, 26, 8, 8, 0},
  {5, 5, 26, 12, 12, 8, 8, 0, 9999, 5, 5, 5, 5, 26, 8, 8, 0},
  {3, 0, 3, 48, 48, 8, 8, 5, 5, 9999, 0, 3, 0, 3, 8, 8, 5},
  {3, 0, 3, 48, 48, 8, 8, 5, 5, 0, 9999, 3, 0, 3, 8, 8, 5},
  {0, 3, 5, 48, 48, 8, 8, 5, 5, 3, 3, 9999, 3, 5, 8, 8, 5},
  {3, 0, 3, 48, 48, 8, 8, 5, 5, 0, 0, 3, 9999, 3, 8, 8, 5},
  {5, 3, 0, 72, 72, 48, 48, 24, 24, 3, 3, 5, 3, 9999, 48, 48, 24},
  {8, 8, 50, 6, 6, 0, 0, 8, 8, 8, 8, 8, 8, 50, 9999, 0, 8},
  {8, 8, 50, 6, 6, 0, 0, 8, 8, 8, 8, 8, 8, 50, 0, 9999, 8},
  {5, 5, 26, 12, 12, 8, 8, 0, 0, 5, 5, 5, 5, 26, 8, 8, 9999}}
  
  fmt.Println(tsp_recocido(caso1,100000,0.01,0.95,10))
}
