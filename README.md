# Online Store

## The reason for the bad reviews

- After successful checkout, most buyers will pay for their order, assuming that the item is being prepared by the seller. In fact, the stock item is no longer available.
- There is no notification to buyers regarding the information of the items they order, especially stock information
- The problem continues, when the order has been canceled but the money they transfer has to wait to be refunded.

## One of the solutions that I offer is

- When the buyer places an order until the checkout process is successful, there is no stock hold which reduces the number of stock items available. Where this stock hold will be deleted when the item is successfully sent to the buyer, when the order is canceled, the stock hold will be deleted and the stock will return according to the amount of stock hold from the canceled one.

- There are still many other solutions that I offer, depending on the capabilities of the system itself and the agreement of business people in the Online Store organization.

## Proof of concept

- the buyer selects the item and fills the quantity to be purchased, and puts it in the basket.
- After checkout, the buyer gets the order_id.
- Then the stock item will decrease automatically in the item collection, and the stock will be recorded into the stock hold collection, based on order_id and item_id.
- Furthermore, the buyer selects a payment method and completes the payment.
- The seller sends the item, and updates the order status to be sent to the buyer. In this process, the stock hold data will be erased and the item stock will actually decrease.
- If there is an order cancellation, be it in the checkout process, payment process or shipping process, the data in the stock hold will also be deleted but the item stock will increase according to the quantity ordered.
- And the most important of all, there is notification to buyers or sellers of every ongoing process.