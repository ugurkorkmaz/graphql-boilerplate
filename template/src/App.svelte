<script lang="ts">
  export let title: string;
  export let description: string;
</script>

<svelte:head>
  <title>{title}</title>
  <meta name="description" content={description} />
</svelte:head>

<main class="w-full h-full flex flex-col justify-center items-center">
    <div class="w-1/2 h-1/2 bg-gray-200 p-4 rounded-lg">
        <h1 class="w-full inline-flex justify-center text-4xl font-bold">{title}</h1>
        <p class="w-full inline-flex justify-center text-xl">{description}</p>

        <div class="w-full flex justify-center mt-8">
            <a href="/playground" class="w-1/2 inline-flex justify-center text-xl font-bold text-white bg-blue-500 rounded-lg p-2">
                Playground
            </a> 
        </div>
    </div>
</main>
