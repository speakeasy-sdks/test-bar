# DrinkInput


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | The name of the drink.                                            |
| `Price`                                                           | *float64*                                                         | :heavy_check_mark:                                                | The price of one unit of the drink in US cents.                   |
| `ProductCode`                                                     | **string*                                                         | :heavy_minus_sign:                                                | The product code of the drink, only available when authenticated. |
| `Type`                                                            | [*shared.DrinkType](../../models/shared/drinktype.md)             | :heavy_minus_sign:                                                | The type of drink.                                                |