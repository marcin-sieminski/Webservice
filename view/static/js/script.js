async function deleteItem(){        
    const itemId = parseInt(document.querySelector(".delete-confirmation").getAttribute("data-item-id"));
    await fetch(`http://localhost/item/delete?id=${itemId}`, { method: "DELETE" })
        .then((response) => {
            if(response.ok){
                location.href = response.url;
            }
        })
}

async function home(){
    location.href = '/';    
}
