<script>
    import image from "../../lib/JoshuaTreeirlpicture.png"
    let text = "Explore the beauty of Joshua Tree National Park.";
    let funFacts = [
        "Joshua Tree National Park is named after the unique Joshua Tree, a species of yucca plant.",
        "The park is home to over 250 species of birds.",
        "Joshua Tree was designated a National Park in 1994.",
        "The park features over 8,000 years of human history, including Native American rock art."
    ];

</script>

<style>
    body {
        margin: 0;
        font-family: 'Arial', sans-serif;
        background-color: #f4f4f4;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        padding: 0;
    }

    .container {
        display: flex;
        align-items: flex-start;
        justify-content: space-between;
        gap: 20px;
        width: 100vw;
        height: 100vh;
        background: white;
        padding: 30px;
        box-sizing: border-box;
        overflow: hidden;
    }

    .colored-box {
        background-color: #87CEEB;
        padding: 40px;
        border-radius: 12px;
        color: #1d4e2c;
        font-size: 1.8em;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: center;
        flex: 1;
        min-width: 400px;
    }

    .logo {
        width: 80px;
        height: 80px;
        border-radius: 50%;
        background-image: url('/JoshuaTree-2.png');
        background-size: cover;
        background-position: center;
        margin-bottom: 20px;
        box-shadow: 0px 2px 6px rgba(0, 0, 0, 0.2);
    }

    .text {
        font-weight: bold;
        font-size: 1.7em;
        margin-bottom: 15px;
    }

    .fun-facts {
        font-size: 1.5em;
        color: #1d4e2c;
    }

    .fun-facts li {
        margin-bottom: 12px;
    }

    .image {
        width: 40vw;
        height: auto;
        border-radius: 8px;
        box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
        margin-left: 20px;
    }

    @media (max-width: 700px) {
        .container {
            flex-direction: column;
            align-items: center;
            justify-content: center;
        }

        .image {
            width: 80vw;
            max-width: 500px;
            margin-top: 20px;
        }

        .colored-box {
            width: 90%;
            padding: 20px;
        }
    }
</style>

<div class="container">
    <div class="colored-box">
        <div class="logo"></div>
        <div class="text">{text}</div>
        <ul class="fun-facts">
            {#each funFacts as fact}
                <li>{fact}</li>
            {/each}
        </ul>
    </div>
    <!-- svelte-ignore a11y_img_redundant_alt -->
    <img class="image" src={image} alt="Joshua Tree Image" />
</div>