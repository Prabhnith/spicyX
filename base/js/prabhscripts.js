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
		} else {
			stickerImg.setAttribute("src", "./img/non-veg-sticker.png");
		}
		stickerDiv.appendChild(stickerImg);
		listItem.appendChild(stickerDiv);
		itemName = document.createElement("div");
		itemName.setAttribute("class", "media-body our-media-body");

		itemHeading = document.createElement("h4");
		itemHeading.setAttribute("class", "media-heading");
		itemHeading.innerHTML = ItemList[i][0];

		addItemSymbol = document.createElement("span");
		addItemSymbol.setAttribute("class", "glyphicon glyphicon-plus add-button pull-right");
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
window.onload = function() {
	loadMenuItems();
}