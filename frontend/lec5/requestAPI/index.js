
function getUser(){
  const result = document.getElementById('result');

  fetch('https://jsonplaceholder.typicode.com/users')
  .then(response => response.json())
  .then(
    data => {
      console.log(data)
      for(user of data){
      let el = document.createElement("tr");
      el.innerHTML = `
        <td>${user.id}</td>
        <td>${user.name}</td>
        <td>${user.email}</td>
        <td>${user.phone}</td>
      `;
      result.append(el);
    }
    }
  )
  .catch((e) => {
    result.innerHTML = `Result error`;
  });

}
