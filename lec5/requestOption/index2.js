
function getUser(){
    const result = document.getElementById('result');
    const body = {
        title: 'foo',
        body: 'bar',
        userId: 1,
    };
    fetch('https://jsonplaceholder.typicode.com/posts',
    { 
        method: "POST",
        body: JSON.stringify({title: 'foo', body: 'bar', userId: 1})
    })
    .then(response => response.json())
    .then(data => {
        
        result.innerHTML = `<p>Result success: ${JSON.stringify(data)}`
    })
    .catch((e) => {
        console.log(2);
        result.innerHTML = `<p>Result error: ${JSON.stringify(e)}`;
    });
}

  
  