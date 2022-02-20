use amiquip::{Connection, Exchange, Publish};
use std::io::{self, Write};
use std::error::Error;

fn main() -> Result<(),Box<dyn Error>> {
    env_logger::init();

    // Open connection.
    let mut connection = Connection::insecure_open("amqp://guest:guest@localhost:5672")?;

    // Open a channel - None says let the library choose the channel ID.
    let channel = connection.open_channel(None)?;

    // Get a handle to the direct exchange on our channel.
    let exchange = Exchange::direct(&channel);

    print!("Message to send over RabbitMQ: ");
    io::stdout().flush().unwrap();
    let mut buffer = String::new();
    io::stdin().read_line(&mut buffer)?;

    // Publish a message to the "hello" queue.
    exchange.publish(Publish::new(buffer.as_bytes(), "rmq_test"))?;

    Ok(connection.close()?)
}
