# Changelog

## [1.5.0](https://github.com/Kavindu-Dodan/go-sdk-contrib/compare/tests/flagd-v1.4.1...tests/flagd/v1.5.0) (2024-04-19)


### 🐛 Bug Fixes

* error handing, add e2e tests ([#334](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/334)) ([dfc3b5a](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/dfc3b5a73e6708aa852a2f2651468de96a754694))
* fix flagd dependencies ([#380](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/380)) ([b7baa69](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/b7baa6990e05f46637917d83b07dbe0f741d0036))
* lint correction on tests ([#276](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/276)) ([b972274](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/b972274655638dd09c90b5974a9f8aca0b04ca13))
* only reset state if necessary ([#126](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/126)) ([51b1a83](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/51b1a83a93311c4b377574176f3f3ed04a7eabb1))
* tidy workspaces ([#97](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/97)) ([c71a5ec](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/c71a5ec7686ec0572bb47f17dbca7e0ec48252d7))


### ✨ New Features

* allow ProviderOptions to be given to scenario initializations ([#113](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/113)) ([3c6c249](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/3c6c249f412c2fb1a841b36007bfd9f57a082eb4))
* domain specific provider for e2e tests ([#443](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/443)) ([b77d0f5](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/b77d0f533cbb93a4f4056c392cd8f4a6b47501c0))
* Eventing support for flagd provider ([#317](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/317)) ([6b373cb](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/6b373cb393729c6f1f2a31b334cf06fac65dd369))
* flagd integration tests module ([#96](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/96)) ([cf69fa4](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/cf69fa4c2214caae34bb9592388b6918e09e451f))
* flagd json eval suite ([#328](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/328)) ([d36f0c5](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/d36f0c50a003a7583bcfdd078219cd46fe7bd77d))
* use symbolic links to bypass flagd receiving empty state in file changes ([#123](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/123)) ([9ad24f1](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/9ad24f10028173ebcbdde161a5aa9c8c8f71a9ed))


### 🧹 Chore

* bump Go to version 1.21 ([#452](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/452)) ([7ec90ce](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/7ec90ce4f9b06670187561afd9e342eed4228be1))
* fix flagd dependencies after mono repo split ([#172](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/172)) ([4b10a18](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/4b10a1833bad5b7f91c6fe2a4c4c2395e14657e4))
* improve e2e test registration ([#400](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/400)) ([b22105c](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/b22105c392e24ce592020a5f1f652547bb5a89e0))
* **main:** release tests/flagd 1.2.2 ([#180](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/180)) ([d215071](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/d215071b1d1714f6fed31bd7163228293bc778a3))
* **main:** release tests/flagd 1.2.3 ([#225](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/225)) ([04184f9](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/04184f9e76f867d17e2ac3791cf87d9d998eb58b))
* **main:** release tests/flagd 1.2.4 ([#277](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/277)) ([fa13423](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/fa1342348c0c8b04a1e6de965657194e4ed9e88a))
* **main:** release tests/flagd 1.3.0 ([#323](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/323)) ([307e3e4](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/307e3e42eaad3690160834b96bcb69e1224d532c))
* **main:** release tests/flagd 1.3.1 ([#335](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/335)) ([cd65569](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/cd655691a2f788ed062167079695aaba99b3d02a))
* **main:** release tests/flagd 1.3.2 ([#381](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/381)) ([cf7fd8c](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/cf7fd8c76df759457b69f3554267990751918ac6))
* **main:** release tests/flagd 1.3.3 ([#403](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/403)) ([910e3bf](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/910e3bf922b4c0290b3e6829332946ad6e12b5aa))
* **main:** release tests/flagd 1.4.0 ([#444](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/444)) ([0811ae8](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/0811ae851181ed6be2660bb6db9eb5c95352fcd1))
* **main:** release tests/flagd 1.4.1 ([#457](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/457)) ([1a8aca1](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/1a8aca1475ff666bd4d768b850aa1425a74ce131))
* release main ([#115](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/115)) ([72d99e4](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/72d99e427d7313897190082731b47e3b093fcf8a))
* release main ([#122](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/122)) ([ffdc02c](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/ffdc02cfcf039a9f243586ba568802e71f5d47ca))
* release main ([#127](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/127)) ([f9b8340](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/f9b8340d0285c23ed0f072666081ad76aba9f018))
* release main ([#94](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/94)) ([b441d7f](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/b441d7fb01e50e5de5b8b6058312817062901f83))
* release main ([#98](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/98)) ([ba789a2](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/ba789a27fc2dd05a19444cb5741a4afe7f061241))
* update flagd e2e tests ([#466](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/466)) ([a8ee306](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/a8ee3068bd3b174bc75a6aeefa0441c61a5b43f7))
* update module github.com/open-feature/go-sdk to v1.4.0 ([#223](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/223)) ([7c8ea46](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/7c8ea46e3e094f746dbf6d80ba6a1b606314e8d7))
* update to go-sdk 1.9.0 ([#404](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/404)) ([11fa3ab](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/11fa3aba065a6dd81caca30e76efc16fb64a25e3))
* upgrade flagd ([#114](https://github.com/Kavindu-Dodan/go-sdk-contrib/issues/114)) ([4965961](https://github.com/Kavindu-Dodan/go-sdk-contrib/commit/49659614d392783cd3434d08f82dbc580f8c839e))

## [1.4.1](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.4.0...tests/flagd/v1.4.1) (2024-03-18)


### 🧹 Chore

* bump Go to version 1.21 ([#452](https://github.com/open-feature/go-sdk-contrib/issues/452)) ([7ec90ce](https://github.com/open-feature/go-sdk-contrib/commit/7ec90ce4f9b06670187561afd9e342eed4228be1))
* update flagd e2e tests ([#466](https://github.com/open-feature/go-sdk-contrib/issues/466)) ([a8ee306](https://github.com/open-feature/go-sdk-contrib/commit/a8ee3068bd3b174bc75a6aeefa0441c61a5b43f7))

## [1.4.0](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.3.3...tests/flagd/v1.4.0) (2024-01-31)


### ✨ New Features

* domain specific provider for e2e tests ([#443](https://github.com/open-feature/go-sdk-contrib/issues/443)) ([b77d0f5](https://github.com/open-feature/go-sdk-contrib/commit/b77d0f533cbb93a4f4056c392cd8f4a6b47501c0))

## [1.3.3](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.3.2...tests/flagd/v1.3.3) (2024-01-23)


### 🧹 Chore

* improve e2e test registration ([#400](https://github.com/open-feature/go-sdk-contrib/issues/400)) ([b22105c](https://github.com/open-feature/go-sdk-contrib/commit/b22105c392e24ce592020a5f1f652547bb5a89e0))
* update to go-sdk 1.9.0 ([#404](https://github.com/open-feature/go-sdk-contrib/issues/404)) ([11fa3ab](https://github.com/open-feature/go-sdk-contrib/commit/11fa3aba065a6dd81caca30e76efc16fb64a25e3))

## [1.3.2](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.3.1...tests/flagd/v1.3.2) (2023-11-14)


### 🐛 Bug Fixes

* fix flagd dependencies ([#380](https://github.com/open-feature/go-sdk-contrib/issues/380)) ([b7baa69](https://github.com/open-feature/go-sdk-contrib/commit/b7baa6990e05f46637917d83b07dbe0f741d0036))

## [1.3.1](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.3.0...tests/flagd/v1.3.1) (2023-09-18)


### 🐛 Bug Fixes

* error handing, add e2e tests ([#334](https://github.com/open-feature/go-sdk-contrib/issues/334)) ([dfc3b5a](https://github.com/open-feature/go-sdk-contrib/commit/dfc3b5a73e6708aa852a2f2651468de96a754694))

## [1.3.0](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.2.4...tests/flagd/v1.3.0) (2023-09-11)


### ✨ New Features

* Eventing support for flagd provider ([#317](https://github.com/open-feature/go-sdk-contrib/issues/317)) ([6b373cb](https://github.com/open-feature/go-sdk-contrib/commit/6b373cb393729c6f1f2a31b334cf06fac65dd369))
* flagd json eval suite ([#328](https://github.com/open-feature/go-sdk-contrib/issues/328)) ([d36f0c5](https://github.com/open-feature/go-sdk-contrib/commit/d36f0c50a003a7583bcfdd078219cd46fe7bd77d))

## [1.2.4](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.2.3...tests/flagd/v1.2.4) (2023-07-21)


### 🐛 Bug Fixes

* lint correction on tests ([#276](https://github.com/open-feature/go-sdk-contrib/issues/276)) ([b972274](https://github.com/open-feature/go-sdk-contrib/commit/b972274655638dd09c90b5974a9f8aca0b04ca13))

## [1.2.3](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.2.2...tests/flagd/v1.2.3) (2023-07-18)


### 🧹 Chore

* update module github.com/open-feature/go-sdk to v1.4.0 ([#223](https://github.com/open-feature/go-sdk-contrib/issues/223)) ([7c8ea46](https://github.com/open-feature/go-sdk-contrib/commit/7c8ea46e3e094f746dbf6d80ba6a1b606314e8d7))

## [1.2.2](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.2.1...tests/flagd/v1.2.2) (2023-05-11)


### 🧹 Chore

* fix flagd dependencies after mono repo split ([#172](https://github.com/open-feature/go-sdk-contrib/issues/172)) ([4b10a18](https://github.com/open-feature/go-sdk-contrib/commit/4b10a1833bad5b7f91c6fe2a4c4c2395e14657e4))

## [1.2.1](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.2.0...tests/flagd/v1.2.1) (2023-02-27)


### Bug Fixes

* only reset state if necessary ([#126](https://github.com/open-feature/go-sdk-contrib/issues/126)) ([51b1a83](https://github.com/open-feature/go-sdk-contrib/commit/51b1a83a93311c4b377574176f3f3ed04a7eabb1))

## [1.2.0](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.1.0...tests/flagd/v1.2.0) (2023-02-24)


### Features

* use symbolic links to bypass flagd receiving empty state in file changes ([#123](https://github.com/open-feature/go-sdk-contrib/issues/123)) ([9ad24f1](https://github.com/open-feature/go-sdk-contrib/commit/9ad24f10028173ebcbdde161a5aa9c8c8f71a9ed))

## [1.1.0](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.0.1...tests/flagd/v1.1.0) (2023-02-13)


### Features

* allow ProviderOptions to be given to scenario initializations ([#113](https://github.com/open-feature/go-sdk-contrib/issues/113)) ([3c6c249](https://github.com/open-feature/go-sdk-contrib/commit/3c6c249f412c2fb1a841b36007bfd9f57a082eb4))

## [1.0.1](https://github.com/open-feature/go-sdk-contrib/compare/tests/flagd/v1.0.0...tests/flagd/v1.0.1) (2023-01-26)


### Bug Fixes

* tidy workspaces ([#97](https://github.com/open-feature/go-sdk-contrib/issues/97)) ([c71a5ec](https://github.com/open-feature/go-sdk-contrib/commit/c71a5ec7686ec0572bb47f17dbca7e0ec48252d7))

## 1.0.0 (2023-01-25)


### Features

* flagd integration tests module ([#96](https://github.com/open-feature/go-sdk-contrib/issues/96)) ([cf69fa4](https://github.com/open-feature/go-sdk-contrib/commit/cf69fa4c2214caae34bb9592388b6918e09e451f))
