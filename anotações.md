## operador curto de declaração, gopher ou marmota

- := parece uma marmota (gopher)
- uso:
  - Tipagem automatica
  - só pode repetir se houverem variaveis novas = atribuição := declaração
  - = atribui
  - := declara
  - A marmota só funciona dentro um code block, por ex:

  func main(){
    x := 10 
  }

  fora deste escopo do code block, ficaria:

  var = 'Bom dia'

  - Terminologia:
    - palavaras reservadas, por ex package, func
    - operadores e operando:
      - sinal é o operador e numeros são os operandos
    - Expressão:
      - é qualquer coisa que produz um resultado.
    - scope
      - Abrangencia

## Palavra chave VAR
  - o Var funciona fora de code block
  - utiliza-se o var somente quando estamos fora do code block

## Tipos
  - São estaticos, tipagem estatica
  - Ao declarar uma variavel do tipo INT por ex, ela só irá aceitar INT
  
- tipos primitivos:
  - Int, String, Bool

- tipos  compostos(são criados pelo usuario):
  - Slice, Array, struct, map.
  {O ato de definir, criar e estruturar tipos compostos chama-se composição.}

## Valor Zero
  - toda variavel se inicia no ambito 0

## O Pacote FMT
  - https://pkg.go.dev/fmt 

## Criando seu proprio tipo
  - são imutaveis

## Conversão, não coerção
  - conversions:
    - para converter um valor para outro tipo, basta colocar o tipo que eu quero
    e o valor em parenteses logo em seguida.