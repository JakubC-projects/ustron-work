<script setup lang="ts">
import Dropdown from './dropdown.vue';
import rulesPDF from '../assets/rules.pdf'

</script>

<template>
    <Dropdown class="mx-5" label="Informacje i reguły">
        <div class="px-8 pb-6">
            <p class="mb-4">
                NOWOŚĆ 2024! Liczymy wpłaty oraz pieniądze wypracowane przez prace, 
                które przekładają się na te cele:<br>
                <ol class="list-decimal list-outside ml-5">
                    <li>Samvirk (mieszkanie w Brunstad)</li>
                    <li>BUK (campy w Brunstad)</li>
                    <li>Utrzymanie zboru lokalnego (prace, 
                        za które trzeba byłoby zapłacić z 
                        budżetu zborowego, ale wykonujemy 
                        je sami, albo przynoszące zysk zborowi)
                    </li>
                </ol>
            </p>
            <a class="mb-4" :href="rulesPDF" target="_blank">LINK do reguł akcji</a>
            <p>
                Niejasności prosimy zgłaszać do liderów zespołów:<br>
                <ul class="list-disc list-outside ml-5">
                    <li>🔵 BLUE: Edwin Nooitgedagt</li>
                    <li>🔴 RED: Przemek Dziadek</li>
                    <li>🟢 GREEN: Sylwia Nooitgedagt</li>
                    <li>🟠 ORANGE: Natalia Bujok</li>
                </ul>
            </p>
        </div>
    </Dropdown>
</template>