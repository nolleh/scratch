# PURPOSING

TEST THAT MAKES ME CONFUSING MAKING INITAL SCRATCH TS PROJECT

# DESCRIPTION 

1. pnpm init
2. tsc --init
3. pnpm add -D ts-node
4. pnpm add -D typescript
5. pnpm add -D @types/node
3. package.json
  - add script (start: ts-node app.ts)
7. target, module: es2022
7. tsconfig.json: baseUrl as '.'
8. tsconfig.json: paths
```json
    "paths": {
      "test/*": [
        "test/*"
      ],
      "*": [
        "src/*" 
      ]
    }
```

## TROUBLE SHOOT
1. cannot find module './app.ts'
  - check baseurl
2. To load ES module, set "type": "module"
  - add package.json
  - it enable ESM in for Node
3. unknown file extension ".ts" for ... app.ts
4. if you use package.jsons module as commonjs, you cannot import es module.(cannot find module)
```typescript

export class Shuffle {
  static shuffle<T>(array: Array<T>, start, length) {
    // ....
  }
}
import { Shuffle } from 'utils/shuffle';

```

5. if you use package.json's type: module was not set.. compile error for 'Cannot use import statement outside a module'

6. unknown file extension:
  --esm



### WORKING OPTION
1. commonjs and remove type:module
run with: ts-node -r tsconfig-paths/register src/app.ts

`package.json`
```json

{
  "name": "single-ts",
  "version": "1.0.0",
  "description": "",
  "main": "src/app.ts",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "ts-node -r tsconfig-paths/register src/app.ts",
    "start2": "ts-node src/app.ts"
  },
  "keywords": [],
  "author": "nolleh",
  "license": "ISC",
  "dependencies": {
    "@types/node": "^18.15.11",
    "ts-node": "^10.9.1"
  },
  "devDependencies": {
    "tsconfig-paths": "^4.2.0",
    "typescript": "^5.0.4"
  }
}
```

`tsconfig.json`

```json

{
  "include": ["src/**/*", "tests/**/*"],
	"exclude": ["node_modules/*", "coverage/*"],
  "compilerOptions": {
    "target": "es2022",
    "module": "commonjs",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "strict": true,
    "skipLibCheck": true,
    "moduleResolution": "node",
    "baseUrl":".",
    "paths": {
      "test/*": [
        "test/*"
      ],
      "*": [
        "src/*" 
      ]
    }

  },
  "ts-node": {
    "esm": true,
    "experimentalSpecifierResolution": "node"
  }
}
```

2. es2022 and add type:module and import statement as relative
and also can you with ts-node src/app.ts

import { Shuffle } from './utils/shuffle';
import { Logger} from './logger';

``package.json``

```json
{
  "name": "single-ts",
  "version": "1.0.0",
  "description": "",
  "main": "src/app.ts",
  "type": "module",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "ts-node -r tsconfig-paths/register src/app.ts",
    "start2": "ts-node src/app.ts"
  },
  "keywords": [],
  "author": "nolleh",
  "license": "ISC",
  "dependencies": {
    "@types/node": "^18.15.11",
    "ts-node": "^10.9.1",
    "tsconfig-paths": "^4.2.0"
  },
  "devDependencies": {
    "typescript": "^5.0.4"
  }
}
```

```json
{
  "include": ["src/**/*", "tests/**/*"],
	"exclude": ["node_modules/*", "coverage/*"],
  "compilerOptions": {
    "target": "es2022",
    "module": "es2022",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "strict": true,
    "skipLibCheck": true,
    "moduleResolution": "node",
    "baseUrl":".",
    "paths": {
      "test/*": [
        "test/*"
      ],
      "*": [
        "src/*" 
      ]
    }

  },
  "ts-node": {
    "esm": true,
    "experimentalSpecifierResolution": "node"
  }
}
```
