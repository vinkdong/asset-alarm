To manager all credit that should very nice

<table>
  <thead>
    <th>api</th>
    <th>method</th>
    <th>body</th>
    <th>use</th>
  </thead>
  <tbody>
   <tr>
    <td>/api/list</td>
    <td>GET/POST</td>
    <td>{}</td>
    <td>Get all credits</td>
   </tr>
   <tr>
     <td>/api/item/add</td>
     <td>POST</td>
     <td>
     <pre lang="javascript">
     {
     	"version":"v0.1",
     	"credit" : {
     		"name":"Vink Bank",
     		"icon":"../icon/vink.logo",
     		"credit":10.000000,
     		"debit":50.000000,
     		"balance":10.000000,
     		"account_date":8,
     		"repayment_date":0
     	}
     }
     </pre>
     </td>
     <td>Add an item to credit</td>
   </tr>
  </tbody>
</table>