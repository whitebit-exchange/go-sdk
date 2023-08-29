v1.1.0 - March 16, 2023

#### Breaking Changes:
- Margin balance response format changed 

#### Highlights
- feature: Add postOnly flag to order responses
- feature: Add ioc flag to order responses
- feature: Add ioc and postOnly parameters to create margin/spot limit
- feature: Add endpoints for get user custom fee and user fee for some market 
- doc: Added an example of using websocket

v1.2.0 - August 28, 2023

#### Highlights
- feature: Add the put_kill_switch endpoint to cancel orders after a specified time (sec) in the "timeout" parameter.
- feature: Add the get_kill_switch_status endpoint retrieves the status of kill-switch timer.
- feature: Add the bulk endpoint creates bulk limit trading orders.
