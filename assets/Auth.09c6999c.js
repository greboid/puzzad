import{S as M,i as N,s as W,a as h,m as q,c as T,o as P,t as D,d as E,f as A,p as F,W as z,ab as H,ac as O,g as S,L as U,e as m,z as v,v as f,x as u,D as k,O as I,n as Y,N as y,A as j,P as B}from"./index.1ca4c1c7.js";import{T as G,F as J,t as C}from"./FlatToast.a7b9e960.js";function K(l){let e,s,a,n,r,i;return{c(){e=m("p"),s=v("You're already logged in, "),a=m("a"),a.textContent="Logout",n=v("?"),f(a,"href","")},m(t,o){T(t,e,o),u(e,s),u(e,a),u(e,n),r||(i=k(a,"click",I(H)),r=!0)},p:Y,d(t){t&&A(e),r=!1,i()}}}function Q(l){let e,s,a,n,r,i,t,o,c,g,$,w,p,_,L;return{c(){e=m("form"),s=m("label"),s.textContent="Email:",a=h(),n=m("input"),r=h(),i=m("label"),i.textContent="Password:",t=h(),o=m("input"),c=h(),g=m("button"),$=v(l[4]),w=h(),p=m("section"),f(s,"for","email"),f(n,"id","email"),f(n,"type","email"),f(n,"name","email"),n.required=!0,f(i,"for","password"),f(o,"id","password"),f(o,"type","password"),f(o,"name","password"),o.required=!0,f(g,"type","submit"),g.disabled=l[6],f(e,"class","basic")},m(d,b){T(d,e,b),u(e,s),u(e,a),u(e,n),y(n,l[1]),u(e,r),u(e,i),u(e,t),u(e,o),y(o,l[2]),u(e,c),u(e,g),u(g,$),u(e,w),u(e,p),p.innerHTML=l[3],_||(L=[k(n,"input",l[8]),k(o,"input",l[9]),k(e,"submit",I(l[0]))],_=!0)},p(d,b){b&2&&n.value!==d[1]&&y(n,d[1]),b&4&&o.value!==d[2]&&y(o,d[2]),b&16&&j($,d[4]),b&8&&(p.innerHTML=d[3])},d(d){d&&A(e),_=!1,B(L)}}}function R(l){let e,s;return e=new J({props:{data:l[10]}}),{c(){q(e.$$.fragment)},m(a,n){P(e,a,n),s=!0},p(a,n){const r={};n&1024&&(r.data=a[10]),e.$set(r)},i(a){s||(D(e.$$.fragment,a),s=!0)},o(a){E(e.$$.fragment,a),s=!1},d(a){F(e,a)}}}function V(l){let e,s,a;function n(t,o){var c;return(c=t[5])!=null&&c.user?K:Q}let r=n(l),i=r(l);return s=new G({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[R,({data:t})=>({10:t}),({data:t})=>t?1024:0]},$$scope:{ctx:l}}}),{c(){i.c(),e=h(),q(s.$$.fragment)},m(t,o){i.m(t,o),T(t,e,o),P(s,t,o),a=!0},p(t,[o]){r===(r=n(t))&&i?i.p(t,o):(i.d(1),i=r(t),i&&(i.c(),i.m(e.parentNode,e)));const c={};o&3072&&(c.$$scope={dirty:o,ctx:t}),s.$set(c)},i(t){a||(D(s.$$.fragment,t),a=!0)},o(t){E(s.$$.fragment,t),a=!1},d(t){i.d(t),t&&A(e),F(s,t)}}}function X(l,e,s){let a;z(l,O,p=>s(5,a=p));let{type:n="login"}=e,r,i,t,o,c;const g=async()=>{let p,_;switch(n){case"signup":({error:_}=await S.auth.signUp({email:r,password:i})),p="An email has been sent to you for verification!";break;case"login":({error:_}=await S.auth.signInWithPassword({email:r,password:i})),p="Login success, redirecting.",await U("/");break}_?C.add({title:"Error",description:_.message,duration:1e4,type:"error"}):C.add({title:"Success",description:p,duration:1e4,type:"success"})};switch(n){case"signup":o="Already have an account? <a href='#/Login'>Login!</a>",c="Sign up";break;case"login":o="Don't have an account? <a href='#/signup'>Sign up!</a>",c="Sign In";break;case"logout":H()}function $(){r=this.value,s(1,r)}function w(){i=this.value,s(2,i)}return l.$$set=p=>{"type"in p&&s(7,n=p.type)},[g,r,i,o,c,a,t,n,$,w]}class ee extends M{constructor(e){super(),N(this,e,X,V,W,{type:7,authAction:0})}get authAction(){return this.$$.ctx[0]}}export{ee as default};
