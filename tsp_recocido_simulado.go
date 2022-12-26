package main
import (
     "fmt" 
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
func tsp_recocido(matriz_distancias [][]int, Temp_Inicial float64, Temp_Final float64, enfriamiento float64, iteraciones int) ([]int,int){
    mejor_costo := math.Inf(1)
    solucion_actual:= solucion_aleatoria(len(matriz_distancias))
    Temp:=Temp_Inicial
    mejor_sol:=solucion_actual
    for Temp>Temp_Final{
        i:=0
        for i< iteraciones{
            sol:=solucion_vecina(solucion_actual)
            if costo(matriz_distancias,sol)<mejor_costo {
                copy(solucion_actual,sol)
                copy(mejor_sol,sol)
                mejor_sol=sol
                mejor_costo=costo(matriz_distancias,sol)
                
            }else{
                var delta float64 = float64(costo(matriz_distancias,sol)-costo(matriz_distancias, solucion_actual))
                r:=rand.Float64()
                if(r<math.Exp(-delta/Temp)){
                   copy(solucion_actual,sol) 
                }
            }
            i++
        }
        Temp=Temp*enfriamiento
    }
    return mejor_sol,int(mejor_costo)
}

func main(){
  start0:=time.Now()
  var tiemposBR17 []time.Duration
  var tiemposFTV35 []time.Duration
  var tiemposFTV44 []time.Duration
  var solucionesBR17 [][]int
  var solucionesFTV35 [][]int
  var solucionesFTV44 [][]int
  var costos_BR17 []int
  var costos_FTV35 []int
  var costos_FTV44 []int

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
  caso2 :=[][]int{{100000000, 26, 82, 65, 102, 100, 147, 134, 69, 117, 125, 38, 40, 13, 38, 31, 22, 103, 143, 94, 104, 123, 140, 98, 58, 38, 30, 67, 120, 149, 100, 141, 93, 162, 62, 66}, 
                  {66, 100000000, 56, 39, 76, 109, 156, 140, 135, 183, 190, 104, 76, 79, 104, 97, 88, 130, 176, 121, 131, 150, 167, 125, 85, 65, 57, 94, 147, 160, 80, 162, 67, 189, 128, 40},
                  {43, 57, 100000000, 16, 20, 53, 100, 84, 107, 155, 168, 81, 53, 56, 81, 74, 65, 146, 186, 137, 147, 166, 183, 141, 101, 81, 73, 110, 163, 164, 102, 166, 71, 205, 105, 62}, 
                  {27, 41, 62, 100000000, 82, 97, 144, 131, 96, 144, 152, 65, 37, 40, 65, 58, 49, 130, 170, 121, 131, 150, 167, 125, 85, 65, 57, 94, 147, 166, 86, 168, 73, 189, 89, 46}, 
                  {106, 120, 63, 79, 100000000, 33, 80, 64, 87, 135, 208, 130, 116, 119, 144, 137, 128, 209, 226, 200, 189, 208, 195, 153, 164, 144, 136, 122, 175, 144, 107, 146, 51, 268, 93, 83}, 
                  {109, 135, 161, 174, 142, 100000000, 47, 34, 54, 102, 175, 97, 97, 96, 128, 135, 131, 198, 193, 203, 213, 232, 249, 207, 167, 147, 139, 176, 229, 222, 204, 224, 148, 235, 60, 175}, 
                  {157, 171, 114, 130, 95, 60, 100000000, 40, 114, 162, 235, 157, 157, 156, 188, 188, 179, 258, 253, 251, 239, 258, 245, 203, 215, 195, 187, 172, 207, 175, 157, 177, 101, 295, 120, 133}, 
                  {143, 169, 132, 148, 113, 34, 31, 100000000, 88, 133, 209, 131, 131, 130, 162, 169, 165, 232, 227, 237, 247, 266, 263, 221, 201, 181, 173, 190, 225, 193, 175, 195, 119, 269, 94, 151}, 
                  {95, 121, 177, 160, 196, 54, 101, 88, 100000000, 48, 158, 83, 83, 82, 114, 121, 117, 184, 179, 189, 199, 218, 235, 193, 153, 133, 125, 162, 215, 244, 195, 236, 188, 221, 46, 161}, 
                  {79, 105, 161, 144, 181, 91, 138, 125, 37, 100000000, 114, 67, 67, 66, 98, 105, 101, 137, 132, 149, 183, 202, 219, 177, 137, 117, 109, 146, 199, 228, 179, 220, 172, 174, 57, 145}, 
                  {204, 230, 286, 269, 306, 216, 255, 237, 162, 125, 100000000, 192, 192, 191, 223, 230, 226, 144, 139, 156, 184, 165, 230, 215, 249, 242, 234, 251, 282, 332, 297, 325, 297, 113, 182, 270}, 
                  {38, 64, 120, 103, 140, 88, 135, 122, 57, 105, 87, 100000000, 28, 25, 31, 38, 47, 110, 105, 122, 142, 161, 178, 136, 96, 76, 68, 105, 158, 187, 138, 179, 131, 147, 50, 104}, 
                  {151, 177, 221, 216, 202, 60, 107, 94, 96, 144, 217, 139, 100000000, 138, 170, 177, 173, 240, 235, 245, 255, 274, 291, 249, 209, 189, 181, 218, 271, 282, 251, 284, 208, 277, 102, 217}, 
                  {13, 39, 95, 78, 115, 87, 134, 121, 56, 104, 112, 25, 27, 100000000, 32, 39, 35, 116, 130, 107, 117, 136, 153, 111, 71, 51, 43, 80, 133, 162, 113, 154, 106, 172, 49, 79}, 
                  {38, 48, 104, 87, 124, 119, 166, 153, 88, 136, 118, 31, 59, 32, 100000000, 7, 16, 123, 136, 114, 124, 143, 160, 118, 78, 58, 50, 87, 140, 169, 120, 161, 115, 178, 81, 88}, 
                  {31, 41, 97, 80, 117, 115, 162, 149, 84, 132, 114, 27, 55, 28, 7, 100000000, 9, 116, 132, 107, 117, 136, 153, 111, 71, 51, 43, 80, 133, 162, 113, 154, 108, 174, 77, 81}, 
                  {22, 32, 88, 71, 108, 122, 169, 156, 91, 139, 123, 36, 62, 35, 16, 9, 100000000, 107, 141, 98, 108, 127, 144, 102, 62, 42, 34, 71, 124, 153, 104, 145, 99, 166, 84, 72}, 
                  {108, 134, 190, 173, 210, 133, 180, 167, 93, 113, 60, 96, 96, 95, 127, 134, 130, 100000000, 46, 63, 116, 135, 162, 147, 166, 146, 138, 175, 221, 257, 208, 249, 201, 120, 86, 174}, 
                  {127, 153, 209, 192, 229, 152, 199, 186, 112, 132, 79, 115, 115, 114, 146, 153, 149, 19, 100000000, 17, 70, 89, 116, 101, 135, 148, 157, 137, 175, 219, 183, 211, 220, 85, 105, 193}, 
                  {153, 179, 235, 218, 255, 178, 225, 212, 138, 158, 105, 141, 141, 140, 172, 179, 175, 45, 57, 100000000, 53, 72, 99, 84, 118, 131, 183, 120, 158, 202, 166, 194, 241, 68, 131, 214}, 
                  {179, 165, 199, 204, 219, 243, 290, 277, 203, 223, 165, 206, 206, 192, 199, 192, 183, 110, 112, 82, 100000000, 19, 46, 31, 65, 78, 149, 67, 105, 149, 113, 141, 188, 95, 196, 161}, 
                  {212, 205, 239, 244, 259, 237, 284, 271, 197, 217, 146, 200, 200, 199, 231, 232, 223, 104, 93, 63, 40, 100000000, 67, 71, 105, 118, 189, 107, 117, 167, 153, 181, 228, 76, 190, 201}, 
                  {184, 170, 204, 209, 224, 248, 295, 282, 208, 228, 175, 211, 211, 197, 204, 197, 188, 115, 150, 106, 72, 57, 100000000, 36, 70, 83, 154, 72, 59, 109, 118, 146, 193, 133, 201, 166}, 
                  {148, 134, 168, 173, 188, 212, 259, 246, 172, 192, 139, 175, 175, 161, 168, 161, 152, 79, 125, 70, 36, 55, 82, 100000000, 34, 47, 118, 36, 89, 118, 82, 110, 157, 131, 165, 130}, 
                  {153, 146, 180, 185, 200, 178, 225, 212, 138, 158, 105, 141, 141, 140, 172, 173, 164, 45, 91, 36, 46, 65, 92, 77, 100000000, 59, 130, 48, 101, 130, 94, 122, 169, 104, 131, 142}, 
                  {173, 166, 200, 205, 220, 198, 245, 232, 158, 178, 125, 161, 161, 160, 192, 193, 184, 65, 111, 56, 66, 85, 112, 97, 20, 100000000, 150, 68, 121, 150, 114, 142, 189, 124, 151, 162}, 
                  {30, 16, 72, 55, 92, 125, 172, 156, 99, 147, 133, 68, 70, 43, 50, 43, 34, 73, 119, 64, 74, 93, 110, 68, 28, 8, 100000000, 37, 90, 119, 70, 111, 83, 132, 92, 56}, 
                  {112, 98, 132, 137, 152, 185, 232, 216, 181, 223, 170, 150, 152, 125, 132, 125, 116, 110, 156, 101, 67, 86, 73, 31, 65, 78, 82, 100000000, 53, 82, 46, 74, 121, 162, 174, 94}, 
                  {144, 130, 164, 169, 184, 217, 256, 225, 213, 261, 234, 182, 184, 157, 164, 157, 148, 174, 209, 165, 131, 116, 59, 95, 129, 122, 114, 93, 100000000, 50, 78, 92, 147, 192, 206, 126}, 
                  {94, 80, 114, 119, 134, 167, 214, 198, 163, 211, 197, 132, 134, 107, 114, 107, 98, 137, 183, 128, 110, 129, 81, 74, 92, 72, 64, 43, 57, 100000000, 28, 56, 103, 196, 156, 76}, 
                  {66, 52, 101, 91, 121, 154, 201, 185, 135, 183, 169, 104, 106, 79, 86, 79, 70, 109, 155, 100, 82, 101, 88, 46, 64, 44, 36, 15, 68, 97, 100000000, 89, 90, 168, 128, 63}, 
                  {129, 124, 86, 102, 106, 139, 186, 170, 193, 241, 241, 167, 139, 142, 158, 151, 142, 181, 227, 172, 154, 173, 160, 118, 136, 116, 108, 87, 140, 168, 72, 100000000, 75, 240, 191, 48}, 
                  {113, 108, 70, 86, 51, 84, 131, 115, 138, 186, 225, 151, 123, 126, 142, 135, 126, 165, 211, 156, 138, 157, 144, 102, 120, 100, 92, 71, 124, 93, 56, 95, 100000000, 224, 144, 32}, 
                  {146, 172, 228, 211, 248, 171, 218, 205, 131, 151, 80, 134, 134, 133, 165, 172, 168, 38, 27, 44, 75, 76, 121, 106, 140, 153, 176, 142, 180, 224, 188, 216, 239, 100000000, 124, 212}, 
                  {102, 128, 184, 167, 203, 61, 108, 95, 7, 55, 165, 90, 90, 89, 121, 128, 124, 191, 186, 196, 206, 225, 242, 200, 160, 140, 132, 169, 222, 251, 202, 243, 195, 228, 100000000, 168}, 
                  {81, 95, 38, 54, 58, 91, 138, 122, 145, 193, 206, 119, 91, 94, 119, 112, 103, 184, 224, 175, 165, 184, 171, 129, 139, 119, 111, 98, 151, 120, 83, 122, 27, 243, 143, 0}}
  caso3:=[][]int{{100000000, 26, 46, 74, 82, 65, 102, 100, 147, 134, 75, 69, 106, 117, 42, 71, 158, 89, 125, 38, 40, 13, 103, 143, 94, 104, 123, 140, 98, 58, 38, 30, 67, 120, 149, 100, 115, 141, 94, 93, 122, 113, 162, 62, 66}, {66, 100000000, 20, 48, 56, 39, 76, 109, 156, 140, 141, 135, 172, 183, 108, 137, 224, 155, 190, 104, 76, 79, 130, 176, 121, 131, 150, 167, 125, 85, 65, 57, 94, 147, 160, 80, 142, 162, 74, 67, 96, 87, 189, 128, 40}, {46, 53, 100000000, 28, 36, 19, 56, 89, 136, 120, 121, 115, 152, 163, 88, 117, 204, 135, 170, 84, 56, 59, 110, 156, 101, 111, 130, 147, 105, 65, 45, 37, 74, 127, 140, 60, 122, 142, 54, 47, 76, 67, 169, 108, 20}, {73, 87, 72, 100000000, 30, 46, 50, 83, 130, 114, 117, 137, 174, 185, 115, 144, 231, 162, 198, 111, 83, 86, 176, 216, 167, 176, 195, 182, 140, 131, 111, 103, 109, 162, 131, 94, 131, 133, 56, 38, 67, 39, 235, 135, 70}, {43, 57, 42, 70, 100000000, 16, 20, 53, 100, 84, 87, 107, 144, 155, 85, 114, 201, 132, 168, 81, 53, 56, 146, 186, 137, 147, 166, 183, 141, 101, 81, 73, 110, 163, 164, 102, 158, 166, 89, 71, 100, 42, 205, 105, 62}, {27, 41, 26, 54, 62, 100000000, 82, 97, 144, 131, 102, 96, 133, 144, 69, 98, 185, 116, 152, 65, 37, 40, 130, 170, 121, 131, 150, 167, 125, 85, 65, 57, 94, 147, 166, 86, 142, 168, 80, 73, 102, 93, 189, 89, 46}, {106, 120, 105, 91, 63, 79, 100000000, 33, 80, 64, 67, 87, 124, 135, 100, 129, 199, 147, 208, 130, 116, 119, 209, 226, 200, 189, 208, 195, 153, 164, 144, 136, 122, 175, 144, 107, 144, 146, 69, 51, 80, 22, 268, 93, 83}, {109, 135, 155, 183, 161, 174, 142, 100000000, 47, 34, 34, 54, 91, 102, 67, 96, 166, 114, 175, 97, 97, 96, 198, 193, 203, 213, 232, 249, 207, 167, 147, 139, 176, 229, 222, 204, 222, 224, 166, 148, 117, 120, 235, 60, 175}, {157, 171, 156, 141, 114, 130, 95, 60, 100000000, 40, 94, 114, 151, 162, 127, 156, 226, 174, 235, 157, 157, 156, 258, 253, 251, 239, 258, 245, 203, 215, 195, 187, 172, 207, 175, 157, 175, 177, 119, 101, 70, 73, 295, 120, 133}, {143, 169, 174, 159, 132, 148, 113, 34, 31, 100000000, 68, 88, 122, 133, 101, 130, 197, 148, 209, 131, 131, 130, 232, 227, 237, 247, 266, 263, 221, 201, 181, 173, 190, 225, 193, 175, 193, 195, 137, 119, 88, 91, 269, 94, 151}, {75, 101, 121, 149, 157, 140, 176, 34, 81, 68, 100000000, 20, 57, 68, 33, 62, 132, 80, 141, 63, 63, 62, 164, 159, 169, 179, 198, 215, 173, 133, 113, 105, 142, 195, 224, 175, 190, 216, 169, 168, 151, 154, 201, 26, 141}, {95, 121, 141, 169, 177, 160, 196, 54, 101, 88, 20, 100000000, 37, 48, 53, 82, 112, 100, 158, 83, 83, 82, 184, 179, 189, 199, 218, 235, 193, 153, 133, 125, 162, 215, 244, 195, 210, 236, 189, 188, 171, 174, 221, 46, 161}, {90, 116, 136, 164, 172, 155, 192, 91, 138, 122, 57, 37, 100000000, 11, 48, 46, 75, 64, 121, 78, 78, 77, 148, 143, 160, 194, 213, 230, 188, 148, 128, 120, 157, 210, 239, 190, 205, 231, 184, 183, 208, 203, 184, 68, 156}, {79, 105, 125, 153, 161, 144, 181, 91, 138, 125, 57, 37, 54, 100000000, 37, 35, 103, 53, 114, 67, 67, 66, 137, 132, 149, 183, 202, 219, 177, 137, 117, 109, 146, 199, 228, 179, 194, 220, 173, 172, 201, 192, 174, 57, 145}, {42, 68, 88, 116, 124, 107, 144, 67, 114, 101, 33, 27, 64, 75, 100000000, 29, 116, 47, 108, 30, 30, 29, 131, 126, 136, 146, 165, 182, 140, 100, 80, 72, 109, 162, 191, 142, 157, 183, 136, 135, 164, 155, 168, 20, 108}, {101, 127, 147, 175, 183, 166, 203, 126, 173, 160, 92, 83, 69, 46, 59, 100000000, 87, 18, 79, 89, 89, 88, 102, 97, 114, 162, 181, 208, 193, 159, 139, 131, 168, 221, 250, 201, 216, 242, 195, 194, 223, 214, 139, 79, 167}, {92, 118, 138, 166, 174, 157, 194, 117, 164, 151, 83, 77, 114, 97, 50, 51, 100000000, 9, 46, 80, 80, 79, 93, 88, 105, 153, 172, 199, 184, 150, 130, 122, 159, 212, 241, 192, 207, 233, 186, 185, 214, 205, 130, 70, 158}, {83, 109, 129, 157, 165, 148, 185, 108, 155, 142, 74, 68, 105, 88, 41, 42, 129, 100000000, 61, 71, 71, 70, 84, 79, 96, 144, 163, 190, 175, 141, 121, 113, 150, 203, 232, 183, 198, 224, 177, 176, 205, 196, 121, 61, 149}, {204, 230, 250, 278, 286, 269, 306, 216, 255, 237, 182, 162, 147, 125, 162, 160, 114, 123, 100000000, 192, 192, 191, 144, 139, 156, 184, 165, 230, 215, 249, 242, 234, 251, 282, 332, 297, 299, 325, 298, 297, 325, 317, 113, 182, 270}, {38, 64, 84, 112, 120, 103, 140, 88, 135, 122, 63, 57, 94, 105, 30, 59, 146, 77, 87, 100000000, 28, 25, 110, 105, 122, 142, 161, 178, 136, 96, 76, 68, 105, 158, 187, 138, 153, 179, 132, 131, 160, 151, 147, 50, 104}, {151, 177, 197, 225, 221, 216, 202, 60, 107, 94, 76, 96, 133, 144, 109, 138, 208, 156, 217, 139, 100000000, 138, 240, 235, 245, 255, 274, 291, 249, 209, 189, 181, 218, 271, 282, 251, 266, 284, 226, 208, 177, 180, 277, 102, 217}, {13, 39, 59, 87, 95, 78, 115, 87, 134, 121, 62, 56, 93, 104, 29, 58, 145, 76, 112, 25, 27, 100000000, 116, 130, 107, 117, 136, 153, 111, 71, 51, 43, 80, 133, 162, 113, 128, 154, 107, 106, 135, 126, 172, 49, 79}, {108, 134, 154, 182, 190, 173, 210, 133, 180, 167, 99, 93, 130, 113, 66, 67, 154, 85, 60, 96, 96, 95, 100000000, 46, 63, 116, 135, 162, 147, 166, 146, 138, 175, 221, 257, 208, 223, 249, 202, 201, 230, 221, 120, 86, 174}, {127, 153, 173, 201, 209, 192, 229, 152, 199, 186, 118, 112, 149, 132, 85, 86, 173, 104, 79, 115, 115, 114, 19, 100000000, 17, 70, 89, 116, 101, 135, 148, 157, 137, 175, 219, 183, 185, 211, 221, 220, 249, 240, 85, 105, 193}, {153, 179, 199, 222, 235, 218, 255, 178, 225, 212, 144, 138, 175, 158, 111, 112, 199, 130, 105, 141, 141, 140, 45, 57, 100000000, 53, 72, 99, 84, 118, 131, 183, 120, 158, 202, 166, 168, 194, 209, 241, 270, 261, 68, 131, 214}, {179, 165, 185, 169, 199, 204, 219, 243, 290, 277, 209, 203, 240, 223, 176, 177, 260, 195, 165, 206, 206, 192, 110, 112, 82, 100000000, 19, 46, 31, 65, 78, 149, 67, 105, 149, 113, 115, 141, 156, 188, 217, 208, 95, 196, 161}, {212, 205, 225, 209, 239, 244, 259, 237, 284, 271, 203, 197, 234, 217, 170, 171, 241, 189, 146, 200, 200, 199, 104, 93, 63, 40, 100000000, 67, 71, 105, 118, 189, 107, 117, 167, 153, 155, 181, 196, 228, 254, 248, 76, 190, 201}, {184, 170, 190, 174, 204, 209, 224, 248, 295, 282, 214, 208, 245, 228, 181, 182, 269, 200, 175, 211, 211, 197, 115, 150, 106, 72, 57, 100000000, 36, 70, 83, 154, 72, 59, 109, 118, 120, 146, 161, 193, 196, 213, 133, 201, 166}, {148, 134, 154, 138, 168, 173, 188, 212, 259, 246, 178, 172, 209, 192, 145, 146, 233, 164, 139, 175, 175, 161, 79, 125, 70, 36, 55, 82, 100000000, 34, 47, 118, 36, 89, 118, 82, 84, 110, 125, 157, 186, 177, 131, 165, 130}, {153, 146, 166, 150, 180, 185, 200, 178, 225, 212, 144, 138, 175, 158, 111, 112, 199, 130, 105, 141, 141, 140, 45, 91, 36, 46, 65, 92, 77, 100000000, 59, 130, 48, 101, 130, 94, 96, 122, 137, 169, 198, 189, 104, 131, 142}, {173, 166, 186, 170, 200, 205, 220, 198, 245, 232, 164, 158, 195, 178, 131, 132, 219, 150, 125, 161, 161, 160, 65, 111, 56, 66, 85, 112, 97, 20, 100000000, 150, 68, 121, 150, 114, 116, 142, 157, 189, 218, 209, 124, 151, 162}, {30, 16, 36, 64, 72, 55, 92, 125, 172, 156, 105, 99, 136, 147, 72, 101, 188, 119, 133, 68, 70, 43, 73, 119, 64, 74, 93, 110, 68, 28, 8, 100000000, 37, 90, 119, 70, 85, 111, 64, 83, 112, 103, 132, 92, 56}, {112, 98, 118, 102, 132, 137, 152, 185, 232, 216, 187, 181, 218, 223, 154, 177, 264, 195, 170, 150, 152, 125, 110, 156, 101, 67, 86, 73, 31, 65, 78, 82, 100000000, 53, 82, 46, 48, 74, 89, 121, 150, 141, 162, 174, 94}, {144, 130, 150, 134, 164, 169, 184, 217, 256, 225, 219, 213, 250, 261, 186, 215, 302, 233, 234, 182, 184, 157, 174, 209, 165, 131, 116, 59, 95, 129, 122, 114, 93, 100000000, 50, 78, 80, 92, 121, 147, 137, 173, 192, 206, 126}, {94, 80, 100, 84, 114, 119, 134, 167, 214, 198, 169, 163, 200, 211, 136, 165, 252, 183, 197, 132, 134, 107, 137, 183, 128, 110, 129, 81, 74, 92, 72, 64, 43, 57, 100000000, 28, 30, 56, 71, 103, 132, 123, 196, 156, 76}, {66, 52, 72, 71, 101, 91, 121, 154, 201, 185, 141, 135, 172, 183, 108, 137, 224, 155, 169, 104, 106, 79, 109, 155, 100, 82, 101, 88, 46, 64, 44, 36, 15, 68, 97, 100000000, 63, 89, 58, 90, 119, 110, 168, 128, 63}, {113, 99, 119, 54, 84, 100, 104, 137, 184, 168, 171, 182, 219, 230, 155, 184, 271, 202, 216, 151, 137, 126, 156, 202, 147, 129, 148, 122, 93, 111, 91, 83, 62, 73, 41, 47, 100000000, 26, 41, 73, 102, 93, 215, 175, 46}, {129, 124, 128, 56, 86, 102, 106, 139, 186, 170, 173, 193, 230, 241, 171, 200, 287, 218, 241, 167, 139, 142, 181, 227, 172, 154, 173, 160, 118, 136, 116, 108, 87, 140, 168, 72, 135, 100000000, 34, 75, 104, 95, 240, 191, 48}, {95, 90, 94, 22, 52, 68, 72, 105, 152, 136, 139, 159, 196, 207, 137, 166, 253, 184, 207, 133, 105, 108, 147, 193, 138, 120, 139, 126, 84, 102, 82, 74, 53, 106, 134, 38, 101, 127, 100000000, 41, 70, 61, 206, 157, 14}, {113, 108, 112, 40, 70, 86, 51, 84, 131, 115, 118, 138, 175, 186, 151, 180, 250, 198, 225, 151, 123, 126, 165, 211, 156, 138, 157, 144, 102, 120, 100, 92, 71, 124, 93, 56, 93, 95, 18, 100000000, 29, 29, 224, 144, 32}, {146, 160, 145, 101, 103, 119, 84, 117, 119, 88, 151, 171, 208, 219, 184, 213, 283, 231, 271, 184, 156, 159, 226, 272, 217, 199, 218, 186, 163, 181, 161, 153, 132, 137, 105, 117, 105, 107, 79, 61, 100000000, 62, 285, 177, 93}, {84, 98, 83, 69, 41, 57, 22, 55, 102, 86, 89, 109, 146, 157, 122, 151, 221, 169, 209, 122, 94, 97, 187, 227, 178, 167, 186, 173, 131, 142, 122, 114, 100, 153, 122, 85, 122, 124, 47, 29, 58, 100000000, 246, 115, 61}, {146, 172, 192, 220, 228, 211, 248, 171, 218, 205, 137, 131, 168, 151, 104, 105, 192, 123, 80, 134, 134, 133, 38, 27, 44, 75, 76, 121, 106, 140, 153, 176, 142, 180, 224, 188, 190, 216, 231, 239, 268, 259, 100000000, 124, 212}, {102, 128, 148, 176, 184, 167, 203, 61, 108, 95, 27, 7, 44, 55, 60, 89, 119, 107, 165, 90, 90, 89, 191, 186, 196, 206, 225, 242, 200, 160, 140, 132, 169, 222, 251, 202, 217, 243, 196, 195, 178, 181, 228, 100000000, 168}, {81, 95, 80, 8, 38, 54, 58, 91, 138, 122, 125, 145, 182, 193, 123, 152, 239, 170, 206, 119, 91, 94, 184, 224, 175, 165, 184, 171, 129, 139, 119, 111, 98, 151, 120, 83, 120, 122, 45, 27, 56, 47, 243, 143, 0}}
   
    for i:=0;i<30;i++{
      start1 := time.Now()
      h,v:=tsp_recocido(caso1,1000000,0.000001,0.725972,1000)
      elapsed1 := time.Since(start1)
      tiemposBR17=append(tiemposBR17,elapsed1)
      solucionesBR17=append(solucionesBR17,h)
      costos_BR17=append(costos_BR17,v)
      start2 :=time.Now()
      j,k:=tsp_recocido(caso2,1000000,0.00001,0.725972,1000)
      elapsed2:=time.Since(start2)
      tiemposFTV35=append(tiemposFTV35,elapsed2)
      solucionesFTV35=append(solucionesFTV35,j)
      costos_FTV35=append(costos_FTV35,k)
      start3:=time.Now()
      l,m:=tsp_recocido(caso3,1000000,0.00001,0.725972,1000)
      elapsed3:=time.Since(start3)
      tiemposFTV44=append(tiemposFTV44,elapsed3)
      solucionesFTV44=append(solucionesFTV44,l)
      costos_FTV44=append(costos_FTV44,m)
  }
  fmt.Println("Los tiempos encontrados para el problema br17: ")
  fmt.Println(tiemposBR17)
  fmt.Println("Las soluciones para el problema br17: ")
  fmt.Println(solucionesBR17)
  fmt.Println("Los costos de cada solución de br17:  ")
  fmt.Println(costos_BR17)
  fmt.Println("Los tiempos encontrados para el problema ftv35: ")
  fmt.Println(tiemposFTV35)
  fmt.Println("Las soluciones para el problema ftv35" )
  fmt.Println(solucionesFTV35)
  fmt.Println("Los costos de cada solución de ftv35: ")
  fmt.Println(costos_FTV35)
  fmt.Println("Los tiempos encontrados para el problema ftv44: ")
  fmt.Println(tiemposFTV44)
  fmt.Println("Las soluciones para el problema ftv44: ")
  fmt.Println(solucionesFTV44)
  fmt.Println("Los costos de cada solucion de ftv44: ")
  fmt.Println(costos_FTV44)
  fmt.Println("El tiempo total de ejecución de este programa es: ", time.Since(start0))
}
