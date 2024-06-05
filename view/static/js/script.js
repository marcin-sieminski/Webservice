async function deleteItem(){        
    const itemId = parseInt(document.querySelector(".delete-confirmation").getAttribute("data-item-id"))
    const response = await fetch(`/item/delete?id=${itemId}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            id: itemId
          })
    })
}

function home(){
    location.href = '/';    
}
