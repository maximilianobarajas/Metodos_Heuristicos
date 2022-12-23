schwef <- function(xx)
{
  
  d <- length(xx)
	
  sum <- sum(xx*sin(sqrt(abs(xx))))

  y <- 418.9829*d - sum
  return(y)
}

#Funcion para soluciones vecinas
solucion_vecina<-function(Sol_actual,desv){
   Q=Sol_actual
   for (y in 1:length(Q)){
      Q[y]=rnorm(1,mean=Q[y],sd=desv)
   }
   return(Q)
}
#Funcion para solucion aleatoria inicial
solucion_aleatoria<-function(d){
   return(runif(n=d, min=400, max=500))
}

#Funcion de recocido
recocido<-function(temp_inicial,temp_final,enfriamiento,iteraciones,dimension){
    mejor_imagen=Inf
    solucion_actual=solucion_aleatoria(dimension)
    Temp=temp_inicial
    mejor_sol=solucion_actual
    while(Temp>temp_final){
        for(i in 1:iteraciones){
            sol=solucion_vecina(solucion_actual,0.1)
            if(schwef(sol)<mejor_imagen){
                solucion_actual=sol
                mejor_sol=sol
                mejor_imagen=schwef(sol)

            }
            else{
                #Criterio de metropolis
                delta=schwef(sol)-schwef(solucion_actual)
                r=runif(1)
                if(r<exp(-delta/Temp)){
                    solucion_actual=sol
            
                }
            }
        }
        Temp=enfriamiento*Temp
    }
    return(c(mejor_imagen,mejor_sol))
}
for(i in 1:30){
h<-recocido(100000,0.1,.95,1000,10)
print("La mejor imagen es: ")
print(h[1])
print("Los valores que generan dicha imagen son: ")
print(h[2:10])
}
