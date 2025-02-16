<script>
    import {app, auth, provider} from '../../firebase_yippee';
    import {signInWithPopup, onAuthStateChanged, GoogleAuthProvider} from 'firebase/auth';
    import {goto} from "$app/navigation";

    let buttonText = "Sign in my scarab";

    function handleLogin() {
        console.log("hi");
        if (!auth.currentUser) {
            signInWithPopup(auth, provider)
                .then(async (result) => {
                    // This gives you a Google Access Token. You can use it to access the Google API.
                    const credential = GoogleAuthProvider.credentialFromResult(result);
                    const token = credential.accessToken;
                    console.log(token);
                    // The signed-in user info.
                    const user = result.user;
                    
                    const idToken = await user.getIdToken();
                    console.log(auth.currentUser,idToken);
                    sendToken(idToken);
                    goto("/");
                }).catch((error) => {
                    console.log("oh naur");
                    // Handle Errors here.
                    const errorCode = error.code;
                    const errorMessage = error.message;
                    console.log(errorCode,errorMessage);
                });
            // redirect(308, "/");           
        } else {
            auth.signOut();
        }
    }

    async function sendToken(token) {
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Bearer "+token);

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            redirect: 'follow'
        };
        const response = await fetch("http://localhost:8080/auth", requestOptions)
            .then(response => response.text())
            .then(result => console.log(result))
            .catch(error => console.log('error', error));
    }
    onAuthStateChanged(auth, (user) => {
        if(user) {
            buttonText = "Log out";
        }
        else {
            buttonText = "Log in";
        }
    })
</script>

<div class="container">
    <button class="signin" on:click={handleLogin}>
        {buttonText}
    </button>
</div>

<style>
    div.container {
        display:flex;
        justify-content: center;
        align-items: center;
        height: 60vh;
    }
    .signin {
        padding: 4em !important;
        font-size: large;
    }
</style>
