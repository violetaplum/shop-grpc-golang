create table product
(
    product_id   varchar(30)                               not null unique primary key comment '상품 아이디',
    product_name varchar(100)                              not null unique comment '상품 이름',
    product_doc  varchar(1000) comment '상품 정보',
    created      timestamp(6) default CURRENT_TIMESTAMP(6) not null,
    modified     timestamp                                 null
);

create table product_stock
(
    product_id varchar(30) not null comment '상품 아이디',
    cnt        int         not null default 0 comment '재고 갯수',
    created    timestamp(6)         default CURRENT_TIMESTAMP(6) not null,
    modified   timestamp   null
);


create table `order`
(
    order_id     varchar(30) not null comment '주문서 아이디',
    user_id      varchar(30) not null comment '주문자 아이디',
    order_status tinyint     not null default 1
        comment '주문 상태 : 1.주문완료 2.배송준비중 3.배송중 4.배송완료 5.주문취소',
    created      timestamp(6)         default CURRENT_TIMESTAMP(6) not null,
    modified     timestamp   null
);

create table order_product
(
    order_product_id varchar(30) not null comment '제품 주문 아이디',
    order_id   varchar(30) not null comment '주문서 아이디',
    product_id varchar(30) not null comment '상품 아이디',
    order_quantity int         not null default 0 comment '주문 수량',
    created      timestamp(6)         default CURRENT_TIMESTAMP(6) not null,
    modified     timestamp   null
);

create table user
(
    user_id   varchar(30)  not null unique primary key comment '사용자 아이디',
    user_name varchar(100) not null comment '사용자 이름',
    is_active tinyint      not null default 1 comment '사용자 상태 : 1.활성 2.비활성',
    created   timestamp(6)          default CURRENT_TIMESTAMP(6) not null,
    modified  timestamp    null
);

show tables;

desc product;