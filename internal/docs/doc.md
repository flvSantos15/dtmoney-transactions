## Pensar em quais entidades vou precisar:

- Cliente (Client):
  o cliente pode ter varias contas.

- Conta (Account):
  Vai realizar as transacoes.

- Transacao (Transaction):
  Vai ser o historico entre as contas.

## Funcionalidades:

- Criar cliente.
- Atualizar cliente.
- Buscar cliente.

- Criar conta para o cliente.
- Atualizar dados da conta do cliente.
- Buscar contas do cliente.

- Realizar deposito.
- Realizar transferencia.
- Realizar saque.

- Buscar historico de transacoes.

## Fluxo:

O front vai enviar as chamadas para um broker.
Essa aplicação vai ler as mensagens e realizar as acoes
Talvez enviar uma resposta para uma possivel notificacao.
