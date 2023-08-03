import{r as ra,b as sa,c as da,aN as ca,h as d,l as r,j as s,L as la,A as n}from"./index.288dc1f4.js";/**
 * tdesign v1.1.17
 * (c) 2023 tdesign
 * @license MIT
 */var oa={bordered:!0,headerBordered:!1,hoverShadow:!1,loading:!1,shadow:!1,size:"medium",theme:"normal"};/**
 * tdesign v1.1.17
 * (c) 2023 tdesign
 * @license MIT
 */var y=ra.exports.forwardRef(function(x,P){var c,i,a=sa(x,oa),l=a.actions,v=a.avatar,S=a.bordered,u=a.children,A=a.className,o=a.cover,h=a.description,f=a.footer,N=a.header,D=a.headerBordered,H=a.hoverShadow,j=a.loading,z=a.shadow,B=a.size,F=a.style,m=a.subtitle,_=a.title,L=a.theme,p=a.status,E=da(),e=E.classPrefix,I=ca(),t=L==="poster2",R=d("".concat(e,"-card"),A,(c={},r(c,I.SIZE.small,B==="small"),r(c,"".concat(e,"-card--bordered"),S),r(c,"".concat(e,"-card--shadow"),z),r(c,"".concat(e,"-card--shadow-hover"),H),c)),b=N||_||m||h||v||l&&!t||p&&t,w=d((i={},r(i,"".concat(e,"-card__header"),b),r(i,"".concat(e,"-card__title--bordered"),D),i)),T=d(r({},"".concat(e,"-card__title"),_)),Z=d(r({},"".concat(e,"-card__subtitle"),m)),C=d(r({},"".concat(e,"-card__actions"),l)),$=d(r({},"".concat(e,"-card__footer"),f)),k=d(r({},"".concat(e,"-card__cover"),o)),q=d(r({},"".concat(e,"-card__avatar"),v)),G=d(r({},"".concat(e,"-card__body"),u)),J=d(r({},"".concat(e,"-card__description"),h)),K=_?s("span",{className:T,children:_}):null,M=m?s("span",{className:Z,children:m}):null,O=h?s("p",{className:J,children:h}):null,Q=v&&s("div",{className:q,children:v}),U=l&&!t&&s("div",{className:C,children:l}),V=l&&t&&s("div",{className:C,children:l}),W=p&&t&&s("div",{className:C,children:p}),X=function(){return N?s("div",{className:w,children:N}):n("div",{className:w,children:[n("div",{className:"".concat(e,"-card__header-wrapper"),children:[Q,n("div",{children:[K,M,O]})]}),U,W]})},Y=o?s("div",{className:k,children:typeof o=="string"?s("img",{src:o,alt:""}):o}):null,aa=u&&s("div",{className:G,children:u}),ea=f&&n("div",{className:$,children:[s("div",{className:"".concat(e,"-card__footer-wrapper"),children:f}),V]}),g=n("div",{ref:P,className:R,style:F,children:[b?X():null,Y,aa,ea]});return j?s(la,{children:g}):g});y.displayName="Card";/**
 * tdesign v1.1.17
 * (c) 2023 tdesign
 * @license MIT
 */var ia=y;export{ia as C};
