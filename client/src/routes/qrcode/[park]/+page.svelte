<script>
    let data = $props();
    let parkname = atob(data.park);
    
    $effect(() => {
        visitPark();
    });

    async function visitPark() {
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Bearer "+token);

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            redirect: 'follow'
        };
        const response = await fetch(`http://localhost:8080/park/${parkname}`, requestOptions)
            .then(response => response.text())
            .then(result => console.log(result))
            .catch(error => console.log('error', error));
    }

</script>

<main>
    <p>{parkname}</p>
</main>