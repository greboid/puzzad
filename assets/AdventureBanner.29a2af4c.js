import{S as q,i as B,s as F,q as S,r as y,h as k,l as z,j as p,u as m,v as D,a as h,m as I,w as g,x as M,b as O,t as Y,d as j,p as E,f as G,y as L,z as U}from"./index.bc9733a8.js";import{A as H}from"./AdventureLogo.6c0be6b4.js";function P(t){let e,n;return{c(){e=k("code"),n=L(t[1]),m(e,"class","svelte-i2ojnk")},m(a,i){h(a,e,i),g(e,n)},p(a,i){i&2&&U(n,a[1])},d(a){a&&j(e)}}}function w(t){let e;return{c(){e=k("div"),e.textContent="ADMIN ONLY",m(e,"class","admin svelte-i2ojnk")},m(n,a){h(n,e,a)},d(n){n&&j(e)}}}function C(t){let e,n=t[4]===0?"Free!":"\xA3"+t[4],a;return{c(){e=k("div"),a=L(n),m(e,"class","price svelte-i2ojnk")},m(i,o){h(i,e,o),g(e,a)},p(i,o){o&16&&n!==(n=i[4]===0?"Free!":"\xA3"+i[4])&&U(a,n)},d(i){i&&j(e)}}}function J(t){let e,n,a,i,o,v,d,b,c;function s(l){t[6](l)}let A={};t[0]!==void 0&&(A.name=t[0]),n=new H({props:A}),S.push(()=>y(n,"name",s));let u=t[1]&&P(t),f=!t[3]&&w(),r=t[4]!==null&&C(t);return{c(){e=k("a"),z(n.$$.fragment),i=p(),u&&u.c(),o=p(),f&&f.c(),v=p(),r&&r.c(),m(e,"class",d="adventurebanner "+t[2].toLowerCase()+" svelte-i2ojnk"),D(e,"background-image","url('"+t[5]+"')"),m(e,"href",b=t[1]?"/#/game/"+t[1]:"/#/adventure/"+t[0])},m(l,_){h(l,e,_),I(n,e,null),g(e,i),u&&u.m(e,null),g(e,o),f&&f.m(e,null),g(e,v),r&&r.m(e,null),c=!0},p(l,[_]){const N={};!a&&_&1&&(a=!0,N.name=l[0],M(()=>a=!1)),n.$set(N),l[1]?u?u.p(l,_):(u=P(l),u.c(),u.m(e,o)):u&&(u.d(1),u=null),l[3]?f&&(f.d(1),f=null):f||(f=w(),f.c(),f.m(e,v)),l[4]!==null?r?r.p(l,_):(r=C(l),r.c(),r.m(e,null)):r&&(r.d(1),r=null),(!c||_&4&&d!==(d="adventurebanner "+l[2].toLowerCase()+" svelte-i2ojnk"))&&m(e,"class",d),(!c||_&3&&b!==(b=l[1]?"/#/game/"+l[1]:"/#/adventure/"+l[0]))&&m(e,"href",b)},i(l){c||(O(n.$$.fragment,l),c=!0)},o(l){Y(n.$$.fragment,l),c=!1},d(l){l&&j(e),E(n),u&&u.d(),f&&f.d(),r&&r.d()}}}function K(t,e,n){let{adventureName:a}=e,{code:i}=e,{status:o}=e,{isPublic:v=!0}=e,{price:d=null}=e,b=a&&G.storage.from("adventures").getPublicUrl(a+"/background.jpg").data.publicUrl;function c(s){a=s,n(0,a)}return t.$$set=s=>{"adventureName"in s&&n(0,a=s.adventureName),"code"in s&&n(1,i=s.code),"status"in s&&n(2,o=s.status),"isPublic"in s&&n(3,v=s.isPublic),"price"in s&&n(4,d=s.price)},[a,i,o,v,d,b,c]}class T extends q{constructor(e){super(),B(this,e,K,J,F,{adventureName:0,code:1,status:2,isPublic:3,price:4})}}export{T as A};
