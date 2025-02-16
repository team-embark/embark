<script>
    import image from "../../lib/JoshuaTreeirlpicture.png";
    import modelImage from "../../lib/joshuatreeafter3.jpg";
    
    let text = "Joshua Tree National Park";
    let funFacts = [
        "Joshua Tree National Park is named after the unique Joshua Tree, a species of yucca plant.",
        "The park is home to over 250 species of birds.",
        "Joshua Tree was designated a National Park in 1994.",
        "The park features over 8,000 years of human history, including Native American rock art."
    ];

    let showModal = false;
    let modalImage = "";
    let modalText = "";

    function openModal() {
        modalImage = modelImage; // Show the second image in the modal
        modalText = "This is how Joshua Tree looks after some time!";
        showModal = true;
    }

    function closeModal() {
        showModal = false;
    }
</script>

<style>
    body {
        margin: 0;
        font-family: 'Arial', sans-serif;
        background-color: #f8f9fa;
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
        padding: 0;
        flex-direction: column;
        text-align: center;
        color: #333;
        overflow: hidden; 
    }

    .container {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 30px;
        width: 100vw; 
        height: 100vh; 
        background: linear-gradient(45deg, #87CEEB, #98FB98);
        border-radius: 15px;
        box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.15);
        padding: 40px;
        box-sizing: border-box;
        flex-wrap: wrap; 
    }

    .content {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        flex: 1;
        padding-right: 20px;
        border-right: 3px solid #fff;
        min-width: 300px;
    }

    .text {
        font-size: 3em; 
        font-weight: bold;
        color: #1d4e2c;
        margin-bottom: 20px;
        text-transform: uppercase;
    }

    .fun-facts {
        font-size: 1.5em; 
        color: #1d4e2c;
        margin-top: 15px;
        line-height: 1.6;
    }

    .fun-facts li {
        margin-bottom: 15px;
        list-style-type: none;
        position: relative;
        padding-left: 20px;
    }

    .fun-facts li::before {
        content: "🌵";
        position: absolute;
        left: 0;
        top: 0;
    }

    .image {
        width: 45vw;
        height: auto;
        border-radius: 12px;
        box-shadow: 0px 6px 12px rgba(0, 0, 0, 0.1);
        transition: transform 0.3s ease;
        cursor: pointer;
    }

    .image:hover {
        transform: scale(1.05);
    }

    /* Modal styling */
    .modal {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background-color: rgba(0, 0, 0, 0.6);
        display: none;
        align-items: center;
        justify-content: center;
        opacity: 0;
        transition: opacity 0.3s ease;
    }

    .modal.show {
        display: flex;
        opacity: 1;
    }

    .modal-content {
        background: white;
        padding: 40px;
        border-radius: 12px;
        max-width: 80%;
        text-align: center;
        position: relative;
    }

    .modal-content img {
        width: 100%;
        height: auto;
        border-radius: 10px;
    }

    .modal-close {
        position: absolute;
        top: 10px;
        right: 10px;
        background: #fff;
        border: none;
        padding: 10px;
        border-radius: 50%;
        font-size: 1.5em;
        cursor: pointer;
        box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.2);
    }

    @media (max-width: 700px) {
        .container {
            flex-direction: column;
            gap: 20px;
            padding: 20px;
        }

        .content {
            width: 100%;
            padding-right: 0;
            border-right: none;
        }

        .image {
            width: 80vw;
            max-width: 600px;
        }
    }
</style>

<div class="container">
    <div class="content">
        <div class="text">{text}</div>
        <ul class="fun-facts">
            {#each funFacts as fact}
                <li>{fact}</li>
            {/each}
        </ul>
    </div>
    <!-- Main Image -->
    <img class="image" src={image} alt="Joshua Tree Image" on:click={openModal} />
</div>

<!-- Modal Popup -->
<div class="modal {showModal ? 'show' : ''}">
    <div class="modal-content">
        <button class="modal-close" on:click={closeModal}>×</button>
        <img src={modalImage} alt="Modal Image" />
        <p>{modalText}</p>
    </div>
</div>