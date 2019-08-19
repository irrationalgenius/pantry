
# Pantry Development Board

### Priority
- [ ] Core Function: ArchiveGuest Implementation - Convert to a Transaction
- [ ] Core Function: ArchiveGuest Implementation - Testing

***********
### Backlog

- [ ] Core Function: GetGuestVisit Controller
- [ ] Core Function: GetGuestVisit Implementation
- [ ] Core Function: UpdateGuestVisit Controller
- [ ] Core Function: RemoveGuestVisit Controller
- [ ] Core Function: RemoveGuestVisit Implementation
- [ ] Feature: Data Locking or User Privs
- [ ] Feature: User Authentication (Login System)
- [ ] Feature: Inventory Tracking


- Update Project Documentation
  - [ ] [Github Markdown](https://guides.github.com/features/mastering-markdown/)
  - [ ] Code Comments
  - [ ] Update Devboard and README

***********
### REST Architecture Compliance
REST? Representational State Transfer, and it is a software architectural style that defines a set of rules to be used when communicating between a client and a server.

1. Endpoints must be Correctly defined
  - [ ] Nouns, not Verbs
  - [ ] Plurality, not singularity
  - [ ] Sub-resource when closely related
2. Implement the Correct HTTP Method
3. Utilize the Correct HTTP Status Codes
4. Document your Endpoints
5. Version your API
6. Send Authentication token by the header instead of the URL
7. Always Validate incoming data

References
- [REST API Best Practices](https://jonathas.com/rest-api-best-practices/)
- [REST: Good Practices for API Design](https://medium.com/hashmapinc/rest-good-practices-for-api-design-881439796dc9)
- [10 Best Practices for Better Restful API](https://blog.mwaysolutions.com/2014/06/05/10-best-practices-for-better-restful-api/)

***********
### [The Twelve-Factor App](https://12factor.net/) Compliance
The Twelve-Factor App is a methodology for building software-as-a-service apps

1. **Codebase**: One codebase tracked in revision control, many deploys
  - [x] VCS: [Git](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control) version 2.20.1.windows.1
  - [x] Github: [Pantry Repository](https://github.com/irrationalgenius/pantry)

2. **Dependencies**: Explicitly declare and isolate dependencies
  - [ ] Gradle

3. **Config**: Store config in the environment
  - [x] [gotenv](https://github.com/subosito/gotenv): A Golang package for storing the initial db connection string
  - [x] postgres: Database stores all other persistent application information. This information is retrieved and set into system vars during launch.

4. **Backing services**: Treat backing services as attached resources
5. **Build, release, run**: Strictly separate build and run stages
6. **Processes**: Execute the app as one or more stateless processes
