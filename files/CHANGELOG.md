# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<a name="Unreleased"></a>
### Unreleased - 2023-08-11

### Fixed

- Add cutting to `bank-payment` for `RecipientName` and `Message` fields [WCOM-6034](https://6kswcz.atlassian.net/browse/WCOM-6034) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Bank connestion error: token expired [WCOM-5991](https://6kswcz.atlassian.net/browse/WCOM-5991) (_[Eduard Balev](eduardbalev@gmail.com))_;
- This week dynamic filter [WCOM-5961](https://6kswcz.atlassian.net/browse/WCOM-5961) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
- Dynamic filters in storage [WCOM-5965](https://6kswcz.atlassian.net/browse/WCOM-5965) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
- Dynamic filters in export documents [WCOM-5964](https://6kswcz.atlassian.net/browse/WCOM-5964) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
- Also allow dynamic filters for custom datetime properties [WCOM-5945](https://6kswcz.atlassian.net/browse/WCOM-5945) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
- Change any translations [WCOM-5918](https://6kswcz.atlassian.net/browse/WCOM-5918) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Targeted communication - role avatar pipe [WCOM-5940](https://6kswcz.atlassian.net/browse/WCOM-5940) (_[Eduard Balev](eduardbalev@gmail.com))_;
### Refactored

- Organization settings - get organization info by VAT number (remove Vies and Ares, refactor endpoint for getting bank account) [WCOM-5993](https://6kswcz.atlassian.net/browse/WCOM-5993) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Replaced old async validation with new one for serie register [WCOM-6038](https://6kswcz.atlassian.net/browse/WCOM-6038) (_[Aleks B](aleks.balev7899@gmail.com))_;
- Mainly refactored template part of dialog. Enhanced experience with async validations [WCOM-3604](https://6kswcz.atlassian.net/browse/WCOM-3604) (_[Aleks B](aleks.balev7899@gmail.com))_;
- Updating Angular to 16.1.5 [WCOM-5838](https://6kswcz.atlassian.net/browse/WCOM-5838) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Replace `ng-intercom` to inhouse `intercom` module [WCOM-5930](https://6kswcz.atlassian.net/browse/WCOM-5930) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Replace `angular2-text-mask` dependency to inhouse `text-mask` module [WCOM-5931](https://6kswcz.atlassian.net/browse/WCOM-5931) (_[Eduard Balev](eduardbalev@gmail.com))_;
### Styled

- Add `line-height` to resolver section for fix [WCOM-6033](https://6kswcz.atlassian.net/browse/WCOM-6033) (_[Eduard Balev](eduardbalev@gmail.com))_;
### Feature

- Enable open chat on document when Rossum section is opened [WCOM-5986](https://6kswcz.atlassian.net/browse/WCOM-5986) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Copy value for document lines by drag dot on field [WCOM-5827](https://6kswcz.atlassian.net/browse/WCOM-5827) (_[Eduard Balev](eduardbalev@gmail.com))_;
- Dialog to choose mode before bulk downloading document files [WCOM-5865](https://6kswcz.atlassian.net/browse/WCOM-5865) (_[Eduard Balev](eduardbalev@gmail.com))_;
- When changing the type of document, also transfer the values [WCOM-5861](https://6kswcz.atlassian.net/browse/WCOM-5861) (_[Eduard Balev](eduardbalev@gmail.com))_;
- New copy workflow method [WCOM-5868](https://6kswcz.atlassian.net/browse/WCOM-5868) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
- Make quick filters dynamic by creating date range right before request [WCOM-5627](https://6kswcz.atlassian.net/browse/WCOM-5627) (_[Krzysztof Kobus](krzysztof.kobus@wflow.com))_;
### Chore

- Remove 'SharedModule' from application dependency [WCOM-6000](https://6kswcz.atlassian.net/browse/WCOM-6000) (_[Eduard Balev](eduardbalev@gmail.com))_;
### CI

- Add and configure the system to generate CHANGELOG.md file [WCOM-5955](https://6kswcz.atlassian.net/browse/WCOM-5955) (_[Eduard Balev](eduardbalev@gmail.com))_;

---

<a name="v4.4.4"></a>
### v4.4.4 - 2023-08-11

### Fixed

- Add cutting to `bank-payment` for `RecipientName` and `Message` fields [WCOM-6034](https://6kswcz.atlassian.net/browse/WCOM-6034) (_[Eduard Balev](eduardbalev@gmail.com))_;

---

<a name="v4.4.3"></a>
### v4.4.3 - 2023-08-01

### Fixed

- Get all roles from all pages for approval path builder [WCOM-5971](https://6kswcz.atlassian.net/browse/WCOM-5971) (_[Eduard Balev](eduardbalev@gmail.com))_;

---

<a name="v4.4.2"></a>
### v4.4.2 - 2023-07-31

### CI

- Up version of dependency list (_[Eduard Balev](eduardbalev@gmail.com))_;
### Performed

- Divide into separate requests to change accounting fields and call one at a time (waterfall). The problem is that a deadlock occurred on the Backend due to simultaneous calls to this endpoint [WCOM-5943](https://6kswcz.atlassian.net/browse/WCOM-5943) (_[Eduard Balev](eduardbalev@gmail.com))_;

---

<a name="v4.4.1"></a>
### v4.4.1 - 2023-07-31

### CI

- Changes in build pipelines (_[Eduard Balev](eduardbalev@gmail.com))_;
### Fixed

- Getting only the first page of roles, there was a problem in the approval path details dialog, fixed by querying to get all roles from all pages [WCOM-5956](https://6kswcz.atlassian.net/browse/WCOM-5956) (_[Eduard Balev](eduardbalev@gmail.com))_;

---

For more details on changes, view the [commit
history](https://bitbucket.org/devteam6k/wflow-main-app/commits/) on Bitbucket.
