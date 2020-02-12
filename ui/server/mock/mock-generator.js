const faker = require('faker');
const moment = require('moment');
const clusters = require('./clusters.mock').getAll().map(x => x.Name);
const nameSpaces = require('./name-spaces.mock').getAll().map(x => x.Name);
const users = require('./users.mock').getAll().map(x => x.Name);
/*
In the server folder
node -e 'console.log(require("./mock/mock-generator.js").getAll(100))' > mock/mock-data.json
* */

const appNames = [...Array(10)].map(() => faker.lorem.slug());
module.exports = {
    getAll(number){
        return [...Array(number)].map(()=>{
            return {
                "Name": faker.random.arrayElement(appNames),
                "Status": faker.random.arrayElement(["successful","failed","running", "timeout"]),
                "Cluster": faker.random.arrayElement(clusters),
                "Namespace": faker.random.arrayElement(nameSpaces),
                "DeployBy": faker.random.arrayElement(users),
                "Time": (new Date(faker.date.between(moment(),moment().subtract(1, 'weeks')))).getTime()
            }
        });
    }
};
