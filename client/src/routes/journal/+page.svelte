<script>
    import image from "../../lib/JoshuaTreeirlpicture.png";
    import modelImage from "../../lib/joshuatreeafter3.jpg";
    import logoImage from "../../lib/JoshuaTree-2.png";  

    let text = "Joshua Tree National Park";
    let funFacts = [
        { fact: "Named after the unique Joshua Tree.", icon: "🌵" },
        { fact: "Home to over 250 species of birds.", icon: "🦅" },
        { fact: "Designated as a National Park in 1994.", icon: "🎉" },
        { fact: "Over 8,000 years of human history, including Native American rock art.", icon: "🖼️" }
    ];

    let showModal = false;
    let modalImage = "";
    let modalText = "";

    function openModal() {
        modalImage = modelImage; 
        modalText = `
            <h2>🔥 The Impact of Wildfires on Joshua Tree National Park 🔥</h2>
            <p><strong>The York Fire (2023)</strong> burned thousands of acres, devastating the fragile ecosystem.</p>
            <ul>
                <li>🌵 Joshua trees are not fire-resistant, and large wildfires threaten their survival.</li>
                <li>🌡 Climate change and human activity increase fire frequency and intensity.</li>
                <li>🚫 Once burned, Joshua trees rarely grow back in affected areas.</li>
                <li>🛑 Conservation efforts, responsible tourism, and fire safety awareness are crucial.</li>
            </ul>
            <p>Protecting our national parks starts with <strong>YOU</strong>.</p>
        `;
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
        justify-content: center;
        align-items: center;
        gap: 30px;
        width: 100vw;
        height: 100vh;
        background: linear-gradient(135deg, #FFB6C1, #FFD700);
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
        min-width: 300px;
        text-align: left;
    }

    .logo {
        width: 200px;
        height: auto;
        margin-bottom: 20px;
    }

    .text {
        font-size: 3em;
        font-weight: bold;
        color: #1d4e2c;
        margin-bottom: 20px;
        text-transform: uppercase;
    }

    .fun-facts {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 20px;
        margin-top: 20px;
        padding: 0;
        list-style: none;
    }

    .fact-card {
        background: #ffffff;
        padding: 20px;
        border-radius: 15px;
        box-shadow: 0px 6px 12px rgba(0, 0, 0, 0.15);
        text-align: center;
        transition: transform 0.3s ease;
    }

    .fact-card:hover {
        transform: translateY(-10px);
    }

    .fact-card .icon {
        font-size: 4em;
        color: #FF6347;
        margin-bottom: 10px;
    }

    .fact-card .fact {
        font-size: 1.2em;
        color: #1d4e2c;
        font-weight: bold;
    }

    .image {
        width: 55vw;
        max-width: 850px;
        height: auto;
        border-radius: 12px;
        box-shadow: 0px 6px 12px rgba(0, 0, 0, 0.1);
        cursor: pointer;
        transition: transform 0.3s ease;
    }

    .image:hover {
        transform: scale(1.05);
    }

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
        background: #fff;
        padding: 30px;
        border-radius: 12px;
        max-width: 800px;
        text-align: left;
        position: relative;
        box-shadow: 0px 8px 20px rgba(0, 0, 0, 0.3);
        font-family: 'Arial', sans-serif;
        line-height: 1.6;
    }

    .modal-content h2 {
        color: #d9534f;
        text-align: center;
        font-size: 1.8em;
        margin-bottom: 15px;
    }

    .modal-content ul {
        padding-left: 20px;
    }

    .modal-content ul li {
        margin-bottom: 10px;
    }

    .modal-content p {
        font-size: 1.1em;
        color: #333;
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
</style>

<div class="container">
    <div class="content">
        <img class="logo" src={logoImage} alt="Logo" /> 
        <div class="text">{text}</div>
        <ul class="fun-facts">
            {#each funFacts as { fact, icon }}
                <li>
                    <div class="fact-card">
                        <div class="icon">{icon}</div>
                        <div class="fact">{fact}</div>
                    </div>
                </li>
            {/each}
        </ul>
    </div>
    <img class="image" src={image} alt="Joshua Tree Image" on:click={openModal} />
</div>

<!-- Modal Popup -->
<div class="modal {showModal ? 'show' : ''}">
    <div class="modal-content">
        <button class="modal-close" on:click={closeModal}>×</button>
        <img src={modalImage} alt="Modal Image" />
        <p>{@html modalText}</p>
    </div>
</div>