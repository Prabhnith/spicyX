
window.onload = function() {
	loadMenuItems();
}

function addItem() {
	var itemCounter = document.getElementById("itemCounter");
	itemCounter.innerHTML = +itemCounter.innerHTML + 1;
}

function removeItem() {
	var itemCounter = document.getElementById("itemCounter");
	if (+itemCounter.innerHTML != 0) {
		itemCounter.innerHTML = +itemCounter.innerHTML - 1;
	}
}

function loadMenuItems() {
	ItemList = [
		["Paneer Shahi and corn Korma", "Mushroom sautéed with baby corn", "270", "V"],
		["Shahi Korma", "onion, tomato, poppy seeds", "230", "N"],
		["Corn Korma", "garlic cloves and spice", "170", "V"],
		["Paneer Shahi and corn Korma", "Mushroom sautéed with baby corn", "270", "V"],
		["Shahi Korma", "onion, tomato, poppy seeds", "230", "N"],
		["Corn Korma", "garlic cloves and spice", "170", "V"]
	];

	for (let i = 0; i < ItemList.length; i++) {
		menuList = document.getElementById("menu-list");

		listItem = document.createElement("li");
		listItem.setAttribute("class", "media");

		stickerDiv = document.createElement("div");
		stickerDiv.setAttribute("class", "media-left");

		stickerImg = document.createElement("img");
		stickerImg.setAttribute("class", "media-object stickers");

		if (ItemList[i][3] == "V") {
			stickerImg.setAttribute("src", "./img/veg-sticker.png");
			stickerImg.setAttribute("alt", "veg");
		} else {
			stickerImg.setAttribute("src", "./img/non-veg-sticker.png");
			stickerImg.setAttribute("alt", "non-veg");
		}
		stickerDiv.appendChild(stickerImg);
		listItem.appendChild(stickerDiv);
		itemName = document.createElement("div");
		itemName.setAttribute("class", "media-body our-media-body");

		itemHeading = document.createElement("h4");
		itemHeading.setAttribute("class", "media-heading");
		if (ItemList[i][3] == "V") {
			itemHeading.setAttribute("role", "veg");
		} else {
			itemHeading.setAttribute("role", "nveg");
		}
		itemHeading.innerHTML = ItemList[i][0];

		addItemSymbol = document.createElement("button");
		addItemSymbol.setAttribute("class", "btn glyphicon glyphicon-plus add-button pull-right");
		addItemSymbol.setAttribute("id", i);
		addItemSymbol.setAttribute("onclick", "addToCart(id);");
		itemHeading.appendChild(addItemSymbol);

		priceTag = document.createElement("p");
		priceTag.setAttribute("class", "pull-right price-tag");
		priceTag.innerHTML = "₹ " + ItemList[i][2];
		itemHeading.appendChild(priceTag);

		itemName.appendChild(itemHeading);

		itemDescription = document.createElement("p");
		itemDescription.setAttribute("class", "our-menu-content");
		itemDescription.innerHTML = ItemList[i][1];
		itemName.appendChild(itemDescription);
		listItem.appendChild(itemName);

		menuList.appendChild(listItem);

		separator = document.createElement("li");
		separator.setAttribute("class", "divider our-divider");
		separator.setAttribute("role", "separator");
		menuList.appendChild(separator);
	}
}

function addToCart(i) {
	var btn = document.getElementById(i);
	var dishName = btn.parentNode.firstChild.data;
	var dishPrice = btn.parentNode.children[1].innerHTML;
	var dishType = btn.parentNode.attributes[1].textContent;
	var item = [];
	item.push(dishName, dishPrice, dishType);
	loadCartItems(item);
};


function loadCartItems(ItemList) {
	console.log(ItemList.length);
	cartItems = document.getElementById("cartItems");
	listItem = document.createElement("li");
	listItem.setAttribute("class", "media");

	plusMinusDiv = document.createElement("div");
	plusMinusDiv.setAttribute("class", "media-left btn-group-vertical cart-plus-minus-buttons-group");

	plusBtn = document.createElement("button");
	plusBtn.setAttribute("class", "btn cart-plus-minus-buttons");
	plusBtn.setAttribute("onclick", "addItem();");

	plusIcon = document.createElement("span");
	plusIcon.setAttribute("class", "glyphicon glyphicon-plus");
	plusBtn.appendChild(plusIcon);
	plusMinusDiv.appendChild(plusBtn);

	itemCounter = document.createElement("span");
	itemCounter.setAttribute("class", "item-counter");
	plusMinusDiv.appendChild(itemCounter);

	minusBtn = document.createElement("button");
	minusBtn.setAttribute("class", "btn cart-plus-minus-buttons");
	minusBtn.setAttribute("onclick", "removeItem();");

	minusIcon = document.createElement("span");
	minusIcon.setAttribute("class", "glyphicon glyphicon-minus");
	minusBtn.appendChild(minusIcon);
	plusMinusDiv.appendChild(minusBtn);

	listItem.appendChild(plusMinusDiv);
	itemName = document.createElement("div");
	itemName.setAttribute("class", "media-body our-media-body");

	itemHeading = document.createElement("h4");
	itemHeading.setAttribute("class", "media-heading");
	itemHeading.innerHTML = ItemList[0];

	stickerImg = document.createElement("img");
	stickerImg.setAttribute("class", "pull-right item-in-cart-sticker");

	if (ItemList[2] == "veg") {
		stickerImg.setAttribute("src", "./img/veg-sticker.png");
	} else {
		stickerImg.setAttribute("src", "./img/non-veg-sticker.png");
	}
	itemHeading.appendChild(stickerImg);

	priceTag = document.createElement("p");
	priceTag.setAttribute("class", "pull-right price-tag item-in-cart-price-tag");
	priceTag.innerHTML = ItemList[1];
	itemHeading.appendChild(priceTag);

	itemName.appendChild(itemHeading);
	listItem.appendChild(itemName);

	cartItems.appendChild(listItem);

	separator = document.createElement("li");
	separator.setAttribute("class", "divider our-divider");
	separator.setAttribute("role", "separator");
	cartItems.appendChild(separator);
}
