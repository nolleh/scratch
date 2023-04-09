mkdir hello-prisma
cd hello-prisma

pnpm init -y
pnpm add -D typescript ts-node @types/node

npx tsc --init

pnpm add -D prisma

npx prisma init --datasource-provider mysql

nsx prisma migrate dev --name init
