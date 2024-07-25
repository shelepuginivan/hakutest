<template>
    <div :class="$style.wrapper">
        <h2 :class="$style.heading">{{ $frontmatter.featuresBefore.title }}</h2>
        <b :class="$style.caption">{{ $frontmatter.featuresBefore.caption }}</b>
    </div>
</template>

<style module>
.wrapper {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    padding-top: 36px;
    padding-bottom: 64px;
}

.heading {
    font-size: 40px;
    font-weight: 700;
}

.caption {
    font-size: 18px;
    color: var(--vp-c-text-2);
}
</style>
